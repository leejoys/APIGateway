package api

import (
	"apigateway/pkg/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Программный интерфейс приложения
type API struct {
	newsDB     storage.IfaceNews
	commentsDB storage.IfaceComments
	r          *mux.Router
}

// Конструктор объекта API
func New(newsDB storage.IfaceNews, commentsDB storage.IfaceComments) *API {
	api := API{}
	api.newsDB = newsDB
	api.commentsDB = commentsDB
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API.
func (api *API) endpoints() {
	//test
	api.r.HandleFunc("/test", api.test).Methods(http.MethodGet)
	//метод вывода списка новостей,
	api.r.HandleFunc("/news/latest", api.latest).Methods(http.MethodGet)
	//метод фильтра новостей,
	api.r.HandleFunc("/news/filter", api.filter).Methods(http.MethodGet)
	//метод получения детальной новости,
	api.r.HandleFunc("/news/detailed", api.detailed).Methods(http.MethodGet)
	//метод добавления комментария.

	//метод получения комментариев по id новости.
	api.r.HandleFunc("/comments/{id}", api.comments).Methods(http.MethodGet)
}

// Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.r
}

//test
func (api *API) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test OK"))
}

//метод вывода списка новостей
//localhost:8080/news/latest?page=1
func (api *API) latest(w http.ResponseWriter, r *http.Request) {
	countS := r.URL.Query().Get("page")
	count, err := strconv.Atoi(countS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := api.newsDB.PostsLatestN(count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

//метод фильтра новостей
//localhost:8080/news/filter?sort=date&direction=desc&count=10&offset=0
func (api *API) filter(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	sort := q.Get("sort")
	direction := q.Get("direction")
	countS := q.Get("count")
	count, err := strconv.Atoi(countS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	offsetS := q.Get("offset")
	offset, err := strconv.Atoi(offsetS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := api.newsDB.PostsByFilter(sort, direction, count, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

//метод получения детальной новости,
//localhost:8080/news/detailed?id=1
func (api *API) detailed(w http.ResponseWriter, r *http.Request) {
	idS := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := api.newsDB.PostsDetailedN(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

//метод получения комментария по id.
//localhost:8080/comments/1
func (api *API) comments(w http.ResponseWriter, r *http.Request) {
	idS := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comments, err := api.commentsDB.Comments(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
