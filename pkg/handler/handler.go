package handler

import (
	_ "EffMob/docs"
	"EffMob/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("/group")
	{
		group.POST("/", h.CreateGroup)           // добавить группу
		group.GET("/", h.GetAllLibrary)          // получить всей библиотеки песен (сортировка происходит по названию группы)
		group.GET("/:id", h.GetAllSongGroupById) // получить все песни группы по id
		group.PATCH("/:id", h.UpdateGroup)       // обновить название группы
		group.DELETE("/:id", h.DeleteGroup)      // удалить группу (и все ее песни)
	}

	song := router.Group("/song")
	{
		song.POST("/", h.CreateSong)      // добавление песни
		song.GET("/", h.GetAllSong)       // получение всех песен
		song.GET("/:id", h.GetSongById)   // получение песни по id
		song.PATCH("/:id", h.UpdateSong)  // обновление песни по id
		song.DELETE("/:id", h.DeleteSong) // удаление песни по id

		verse := song.Group("/:id/verse/:verse")
		{
			verse.GET("", h.GetVerses) // получение 1 или более куплетов с параметром limit который указывает на колличество куплетов которое будет получено после указанного
		}
	}

	return router
}
