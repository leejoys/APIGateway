/*Реализовать скелет сервиса API Gateway.
Создайте новый проект в IDE, назовите его APIGateway.
Нужно в рамках этого проекта поднять HTTP-сервер для адреса:

http://localhost:8080/

И добавить следующие обработчики:

метод вывода списка новостей,
метод фильтра новостей,
метод получения детальной новости,
метод добавления комментария.

Пока нам не нужно делать полную реализацию. Вам требуется:
*добавить модели NewsFullDetailed, NewsShortDetailed, Comment
*в методах, которые отдают данные, в теле определить объект или массив
(в зависимости от метода) и возвращать эти данные в качестве ответа

То есть вам нужно использовать то, что называется hard-code.
Методы будут возвращать всегда одно и то же, но пока это не важно.
*/

package main

import (
	"apigateway/pkg/api"
	"apigateway/pkg/storage/memdbnews"
	"apigateway/pkg/storage/mongocomm"
	"fmt"
	"log"
	"net/http"
	"os"
)

type server struct {
	api *api.API
}

func main() {
	srv := server{}
	newsDB := memdbnews.New()

	// Создаём объект базы данных MongoDB.
	pwd := os.Getenv("Cloud0pass")
	connstr := fmt.Sprintf(
		"mongodb+srv://sup:%s@cloud0.wspoq.mongodb.net/gonews?retryWrites=true&w=majority",
		pwd)
	commentsDB, err := mongocomm.New("comments", connstr)
	if err != nil {
		log.Fatalf("mongocomm.New error:%s", err)
	}
	srv.api = api.New(newsDB, commentsDB)
	log.Fatal(http.ListenAndServe("localhost:8080", srv.api.Router()))
}
