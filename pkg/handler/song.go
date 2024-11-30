package handler

import (
	"EffMob/logger"
	"EffMob/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
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

	songInfo, err := GetSongInfo(viper.GetString("apiUrl"), input.GroupName, input.SongName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to fetch song info: "+err.Error())
		logger.Log.Error("error fetching song info", err.Error())
		return
	}

	id, err := h.service.CreateSong(input.GroupName, input.SongName, songInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to add song in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"` // Дата выхода песни
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (h *Handler) ExternalApi(c *gin.Context) {

	response := SongDetail{
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllSong(c *gin.Context) {
	songs, err := h.service.GetAllSongs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get all songs", err.Error())
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *Handler) GetSongById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get id", err.Error())
		return
	}

	song, err := h.service.Song.GetSongById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get song by id", err.Error())
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *Handler) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get id", err.Error())
		return
	}

	var input models.UpdateSong
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input in updateSong", err.Error())
		return
	}

	err = h.service.Song.UpdateSong(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to update song in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeleteSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get id", err.Error())
		return
	}

	err = h.service.Song.DeleteSong(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to delete song in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
