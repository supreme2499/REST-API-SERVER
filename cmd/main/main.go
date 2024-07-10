package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"rest-api-server/internal/config"
	"rest-api-server/internal/user"
	"rest-api-server/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	cfg := config.GetConfig()

	//передаём написанный наш логер в функцию меин
	logger := logging.GetLogger()

	//мультиплекстор или роутер. его мы используем для выполнения наших http запросов
	logger.Info("Создание роутера")
	router := httprouter.New()

	//вызываем функцию который возвращает ссылку на структуру handler которая нам нужна
	//для работы наших методов обработки событий. Тоесть ты передаём нашу структуру в
	//наш мэин что бы мы могли им пользоваться здесь
	handler := user.NewHandler(logger)

	//вызываем метод Register используя нашу структуру, которую мы передали ранее
	//а в метод мы передаём роутер. Метод регистер это наш обработчик событий
	//окоторый отвечает за ответы на запросы к серверу
	handler.Register(router)

	//вызов функции которая запускает наш сервер
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	//передаём логер в функцию
	logger := logging.GetLogger()

	//config
	//делаем проверку на то где мы запускаем наш сервер. на порту или на сокете
	//честно пока что я не знаю как запускать сервер на сокете и запуск на сокете я
	// просто бездумне переписал, но скоро я это исправлю. как запускать на порту я
	//знаю

	//оюбъявляем переменные Листен и листенерр так как они оибе используются несколько раз
	//и лучше их вынести отдельно
	var listener net.Listener
	var ListenErr error

	//config
	//делаем проверку на то где мы запускаем наш сервер. на порту или на сокете
	//честно пока что я не знаю как запускать сервер на сокете и запуск на сокете я
	// просто бездумне переписал, но скоро я это исправлю. как запускать на порту я
	//знаю

	//БАГ
	//пока что при запуске сервера на сокете путь криво записывается(там почему-то двойные слеши)

	//мы делаем проверку конфига на то где мы запускаем сервер на тсп или на сокете
	//если на сокете, то мы создаём путь для него и там запускаем
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("создание листинга сокета")
		socketPath := filepath.Join(appDir, "app.sock")

		logger.Info("создание юниксов")
		listener, ListenErr = net.Listen("unix", socketPath)
		logger.Infof("сервер запущен на сокете %s", socketPath)

		//в остальных случаях мы запускаем на тсп и передаём параметры из конфига
	} else {
		logger.Info("создание листинга tcp")
		listener, ListenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("сервер запущен на %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	//так же мы отдельно вынесли обработчик событий т. к. нам проще просто в конце один раз проверить на ошибку
	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}
	server := &http.Server{
		//хендлер(обработчик событий)
		Handler: router,
		//время ожидания на запись
		WriteTimeout: 15 * time.Second,
		//время ожидания на чтение
		ReadTimeout: 15 * time.Second,
	}
	//запуск сервера + если появится ошибка код выйдет в лог фатал
	logger.Fatal(server.Serve(listener))
}
