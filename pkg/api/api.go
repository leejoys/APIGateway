package api

import (
	"apigateway/pkg/storage"
	"apigateway/pkg/storage/memdb"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Программный интерфейс приложения
type API struct {
	db storage.Interface
	r  *mux.Router
}

// Конструктор объекта API
func New() *API {
	api := API{}
	api.db = memdb.New()
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API.
func (api *API) endpoints() {
	//метод вывода списка новостей,
	api.r.HandleFunc("/news/latest?page={n}", api.latest).Methods(http.MethodGet)
	//метод фильтра новостей,
	//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
	//метод получения детальной новости,
	//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
	//метод добавления комментария.
	//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
}

// Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.r
}

//метод вывода списка новостей,
func (api *API) latest(w http.ResponseWriter, r *http.Request) {
	ns := mux.Vars(r)["n"]
	n, err := strconv.Atoi(ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := api.db.PostsN(n)
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

//метод фильтра новостей,
//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
//метод получения детальной новости,
//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
//метод добавления комментария.
//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
