package handler

import (
	"EffMob/logger"
	"EffMob/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) GetAllLibrary(c *gin.Context) {
	library, err := h.service.Group.GetAllLibrary()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get library from repository")
		return
	}

	c.JSON(http.StatusOK, library)
}

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
