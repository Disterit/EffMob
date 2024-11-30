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

// @Summary CreateSong
// @Tags Song
// @Description Create a song with request to an external API and save enriched data to the database
// @Accept json
// @Produce json
// @Param input body AddSong true "Song information (group and song)"
// @Success 200 {object} map[string]interface{} "Song created successfully"
// @Failure 400 {object} errorResponse "Invalid input data"
// @Failure 500 {object} errorResponse "Internal server error"
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

// @Summary GetSongById
// @Tags Song
// @Description Get a song by its ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 400 {object} errorResponse "Invalid ID provided"
// @Failure 404 {object} errorResponse "Song not found"
// @Failure 500 {object} errorResponse "Internal server error"
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

// @Summary GetSongById
// @Tags Song
// @Description Get detailed information about a song by its ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song "Detailed song information"
// @Failure 400 {object} errorResponse "Invalid ID format"
// @Failure 500 {object} errorResponse "Internal server error"
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

// @Summary UpdateSong
// @Tags Song
// @Description Update information about an existing song
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param input body models.UpdateSong true "Updated song data"
// @Success 200 {object} StatusResponse "Status of the update operation"
// @Failure 400 {object} errorResponse "Invalid ID format or bad request data"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /song/{id} [patch]

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

// @Summary DeleteSong
// @Tags Song
// @Description Delete a song by its ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} StatusResponse "Status of the delete operation"
// @Failure 400 {object} errorResponse "Invalid ID format"
// @Failure 500 {object} errorResponse "Internal server error"
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
