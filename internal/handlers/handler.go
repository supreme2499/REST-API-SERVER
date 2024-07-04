package handlers

import "github.com/julienschmidt/httprouter"

//интерфейс описывающий наш метод-обработчик событий. лично я вообще не понимаю зачем он здесь
//он вроде нигде не использется, ну хз пусть пока что будет
type Handler interface {
	Register(router *httprouter.Router)
}
