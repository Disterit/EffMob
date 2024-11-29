package handler

import (
	"EffMob/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetVerses(c *gin.Context) {
	songId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get id song")
		return
	}

	verseId, err := strconv.Atoi(c.Param("verse"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get verse song")
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get limit song or not valid range")
		return
	}

	verses, err := h.service.Verse.GetVerses(songId, verseId, limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get verses")
		return
	}

	c.JSON(http.StatusOK, verses)
}

func (h *Handler) UpdateVerse(c *gin.Context) {}
