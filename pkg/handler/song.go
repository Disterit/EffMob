package handler

import (
	"EffMob/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddSong struct {
	GroupName string `json:"group" binding:"required"`
	SongName  string `json:"song" binding:"required"`
}

func (h *Handler) CreateSong(c *gin.Context) {
	var input AddSong
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input in createSong", err.Error())
		return
	}

	id, err := h.service.CreateSong(input.GroupName, input.SongName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to add song in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetSongById(c *gin.Context) {}

func (h *Handler) GetAllSong(c *gin.Context) {}

func (h *Handler) UpdateSong(c *gin.Context) {}

func (h *Handler) DeleteSong(c *gin.Context) {}
