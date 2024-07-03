package main

import (
	"log"
	"net"
	"net/http"
	"rest-api-server/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//мультиплекстор ил роутер. его мы используем для выполнения наших http запросов
	log.Println("Создание роутера")
	router := httprouter.New()
	handler := user.NewHandler()

	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	//что-то типо конфига куда мы передаём тип соединения(network) и порт на котором будет работать наш сервер
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Ошибка конфига", err)
	}

	server := &http.Server{
		//хендлер(обработчик событий)
		Handler: router,
		//время ожидания на запись
		WriteTimeout: 15 * time.Second,
		//время ожидания на чтение
		ReadTimeout: 15 * time.Second,
	}
	//запуск сервера
	log.Fatal(server.Serve(listener))
}
