package user

import (
	"net/http"
	"rest-api-server/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/user"
	userURL  = "/user/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartaillyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

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
