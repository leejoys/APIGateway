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
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Программный интерфейс сервера GoNews
type API struct {
	r *mux.Router
}

// Конструктор объекта API
func New() *API {
	api := API{}
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API.
func (api *API) endpoints() {
	// получить n последних новостей
	//api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
	// веб-приложение
	//api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

// Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.r
}

func main() {
	api := New()
	log.Fatal(http.ListenAndServe("localhost:8080", api.Router))
}
