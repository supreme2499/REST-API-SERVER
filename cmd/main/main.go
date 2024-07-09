package main

import (
	"log"
	"net"
	"net/http"
	"rest-api-server/internal/user"
	"rest-api-server/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	logger := logging.GetLogger()

	//мультиплекстор или роутер. его мы используем для выполнения наших http запросов
	logger.Info("Создание роутера")
	router := httprouter.New()

	//вызываем функцию который возвращает ссылку на структуру handler которая нам нужна
	//для работы наших методов обработки событий. Тоесть ты передаём нашу структуру в
	//наш мэин что бы мы могли им пользоваться здесь
	handler := user.NewHandler()

	//вызываем метод Register используя нашу структуру, которую мы передали ранее
	//а в метод мы передаём роутер. Метод регистер это наш обработчик событий
	//окоторый отвечает за ответы на запросы к серверу
	handler.Register(router)

	//вызов функции которая запускает наш сервер
	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	//что-то типо конфига куда мы передаём тип соединения(network) и порт на котором будет работать наш сервер
	logger.Info("создание листинга")
	listener, err := net.Listen("tcp", ":8080")

	//проверяем наличие ошибки
	if err != nil {
		log.Fatal("listener is error", err)
	}
	logger.Info("Запуск обработчика событий")
	server := &http.Server{
		//хендлер(обработчик событий)
		Handler: router,
		//время ожидания на запись
		WriteTimeout: 15 * time.Second,
		//время ожидания на чтение
		ReadTimeout: 15 * time.Second,
	}
	//запуск сервера + если появится ошибка код выйдет в лог фатал
	logger.Info("сервер запущен")
	log.Fatal(server.Serve(listener))
}
