package handler

import (
	"EffMob/logger"
	"EffMob/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateGroup godoc
// @Summary CreateGroup
// @Tags Group
// @Description Create a new group
// @Accept json
// @Produce json
// @Param input body models.Group true "Group details"
// @Success 200 {object} map[string]interface{} "ID of the created group"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /group [post]
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

// GetAllLibrary godoc
// @Summary GetAllLibrary
// @Tags Group
// @Description Get all groups in the library
// @Produce json
// @Success 200 {object} map[string][]models.Song "map of group"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /group [get]
func (h *Handler) GetAllLibrary(c *gin.Context) {
	library, err := h.service.Group.GetAllLibrary()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get library from repository")
		return
	}

	c.JSON(http.StatusOK, library)
}

// GetAllSongGroupById godoc
// @Summary GetAllSongGroupById
// @Tags Group
// @Description Get all songs in a group by group ID
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {array} map[string][]models.Song "Group and they song"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /group/{id} [get]
func (h *Handler) GetAllSongGroupById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get song group id", err.Error())
		return
	}

	songs, err := h.service.Group.GetAllSongGroupById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get song group by id", err.Error())
		return
	}

	c.JSON(http.StatusOK, songs)
}

// UpdateGroup godoc
// @Summary UpdateGroup
// @Tags Group
// @Description Update an existing group
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param input body models.Group true "Updated group details"
// @Success 200 {object} StatusResponse "Status of the update operation"
// @Failure 400 {object} ErrorResponse "Invalid input or group not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /group/{id} [patch]
func (h *Handler) UpdateGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to get group id", err.Error())
		return
	}

	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input group", err.Error())
		return
	}

	err = h.service.Group.UpdateGroup(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to update group in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// DeleteGroup godoc
// @Summary DeleteGroup
// @Tags Group
// @Description Delete a group by its ID
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} StatusResponse "Status of the delete operation"
// @Failure 400 {object} ErrorResponse "Invalid ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /group/{id} [delete]
func (h *Handler) DeleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to delete group id", err.Error())
		return
	}

	err = h.service.Group.DeleteGroup(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to delete group in repository", err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
