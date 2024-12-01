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

// CreateSong godoc
// @Summary Create a new song
// @Tags Song
// @Description Add a new song to the database
// @Accept json
// @Produce json
// @Param input body AddSong true "Song data"
// @Success 200 {object} map[string]interface{} "Song created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /song [post]
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

// GetAllSong godoc
// @Summary Get all songs
// @Tags Song
// @Description Retrieve a list of all songs
// @Produce json
// @Success 200 {object} []models.Song "List of songs"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /song [get]
func (h *Handler) GetAllSong(c *gin.Context) {
	songs, err := h.service.GetAllSongs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get all songs", err.Error())
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetSongById godoc
// @Summary Get song by ID
// @Tags Song
// @Description Retrieve a song by its ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song "Song details"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /song/{id} [get]
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

// UpdateSong godoc
// @Summary Update song by ID
// @Tags Song
// @Description Update an existing song by its ID
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param input body models.UpdateSong true "Updated song data"
// @Success 200 {object} StatusResponse "Update status"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /song/{id} [patch]
func (h *Handler) UpdateSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get id", err.Error())
		return
	}

	var input models.UpdateSong
	if err = c.ShouldBindJSON(&input); err != nil {
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

// DeleteSong godoc
// @Summary Delete song by ID
// @Tags Song
// @Description Delete a song from the database by its ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} StatusResponse "Delete status"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /song/{id} [delete]
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
