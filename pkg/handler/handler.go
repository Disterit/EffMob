package handler

import (
	"EffMob/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	song := router.Group("/song")
	{
		song.POST("")       // добавление песни
		song.GET("")        // получение всех песен
		song.GET("/:id")    // получение песни по id
		song.PUT("/:id")    // обновление песни по id
		song.DELETE("/:id") // удаление песни по id

		verse := song.Group("/:id/verse/:verse")
		{
			verse.GET("") // получение 1 или более куплетов с параметром limit который указывает на колличество куплетов которое будет получено после указанного
			verse.PUT("") //обновление куплета по id
		}
	}

	return router
}
