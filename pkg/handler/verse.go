package handler

import (
	"EffMob/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetVerses godoc
// @Summary GetVerses
// @Tags Song
// @Description Get verses for a song by song ID and verse ID with optional limit
// @Produce json
// @Param id path int true "Song ID"
// @Param verse path int true "Verse ID"
// @Param limit query int false "Limit for the number of verses to fetch"
// @Success 200 {array} map[string]string "List of verses"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /song/{id}/verse/{verse} [get]
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
