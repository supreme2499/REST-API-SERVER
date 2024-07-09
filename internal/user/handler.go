package user

import (
	"net/http"
	"rest-api-server/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

// константы - пути по которым отправляются запросы
const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

// вау пустая структура я хз зачем она нужна, ну ладно....
// пусть пока что будет :(
type handler struct {
}

// функция которая возвращает интерфейс(возвращает ссылку на пустую структуру описанную чуть выше)
// тоже не понимаю зачем.....
func NewHandler() handlers.Handler {
	return &handler{}
}

// обработчик событий, регистер - регистрирует запросы, в методе просто перечисленны все
// возможные запросы которые написанны ниже
func (h *handler) Register(router *httprouter.Router) {

	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartaillyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

// пока что в запросах просто затычки
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями1"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями2"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями3"))
}

func (h *handler) PartaillyUpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями4"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями5"))
}
