package user

import (
	"fmt"
	"net/http"
)

const (
	usersURL = "/user"
	userURL  = "/user/:uuid"
)

type handler struct {
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATH(userURL, h.PartaillyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("это лист с пользователями"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, id httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", id.ByName("name"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, id httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, id httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *handler) PartaillyUpdateUser(w http.ResponseWriter, r *http.Request, id httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
