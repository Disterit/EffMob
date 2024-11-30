package handler

import (
	"EffMob/logger"
	"EffMob/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input group")
		return
	}

	id, err := h.service.Group.CreateGroup(input.GroupName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to create group in repository")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type OutputLibrary struct {
	Group models.Group  `json:"group"`
	Songs []models.Song `json:"songs"`
}

func (h *Handler) GetAllLibrary(c *gin.Context) {
	library, err := h.service.Group.GetAllLibrary()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get library from repository")
		return
	}

	c.JSON(http.StatusOK, library)
}

func (h *Handler) GetAllSongGroupById(c *gin.Context) {}

func (h *Handler) UpdateGroup(c *gin.Context) {}

func (h *Handler) DeleteGroup(c *gin.Context) {}
