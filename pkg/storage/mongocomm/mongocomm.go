package mongocomm

import (
	"apigateway/pkg/storage"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrorDuplicatePost error = errors.New("MongoDB E11000")

// Хранилище данных.
type Store struct {
	c  *mongo.Client
	db *mongo.Database
}

//New - Конструктор объекта хранилища.
func New(name string, connstr string) (*Store, error) {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(connstr))
	if err != nil {
		return nil, err
	}
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}

	s := &Store{c: client, db: client.Database(name)}
	t := true
	_, err = s.db.Collection("comments").Indexes().CreateOne(
		context.Background(), mongo.IndexModel{
			Keys:    bson.D{{Key: "title", Value: 1}},
			Options: &options.IndexOptions{Unique: &t}})
	if err != nil {
		s.c.Disconnect(context.Background())
		return nil, err
	}

	return s, nil
}

//Close - освобождение ресурса
func (s *Store) Close() {
	s.c.Disconnect(context.Background())
}

func (s *Store) DropDB() error {
	return s.db.Drop(context.Background())
}

//Comments - получение всех комментариев
func (s *Store) Comments(id int) ([]storage.Comment, error) {

	coll := s.db.Collection("comments")
	ctx := context.Background()
	filter := bson.D{}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	comments := []storage.Comment{}
	for cur.Next(ctx) {
		var c storage.Comment
		err = cur.Decode(&c)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

//CommentsN - получение n последних комментариев
func (s *Store) PostsN(n int) ([]storage.Comment, error) {

	coll := s.db.Collection("comments")
	ctx := context.Background()
	options := options.Find()
	options.SetLimit(int64(n))
	options.SetSort(bson.D{{Key: "$natural", Value: -1}})
	filter := bson.D{}
	cur, err := coll.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	comments := []storage.Comment{}
	for cur.Next(ctx) {
		var c storage.Comment
		err = cur.Decode(&c)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

//AddComment - создание нового комментария
func (s *Store) AddComment(c storage.Comment) error {
	coll := s.db.Collection("comments")
	_, err := coll.InsertOne(context.Background(), c)

	if mongo.IsDuplicateKeyError(err) {
		return ErrorDuplicatePost
	}
	return err
}

//UpdateComment - обновление по id значения idparent, content, idchild и idnews
func (s *Store) UpdateComment(c storage.Comment) error {
	coll := s.db.Collection("comments")
	filter := bson.D{{Key: "id", Value: c.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "idparent", Value: c.ParentID},
		{Key: "content", Value: c.Content},
		{Key: "idchild", Value: c.ChildsIDs},
		{Key: "idnews", Value: c.IDNews}}}}
	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

//DeleteComment - удаляет комментарий по id
func (s *Store) DeleteComment(c storage.Comment) error {
	coll := s.db.Collection("comments")
	filter := bson.D{{Key: "id", Value: c.ID}}
	_, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
