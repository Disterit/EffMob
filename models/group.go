package models

type Group struct {
	Id        int    `json:"id"`
	GroupName string `json:"name" binding:"required"`
}
