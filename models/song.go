package models

import "time"

type Song struct {
	Id          int       `json:"id"`
	GroupId     int       `json:"group_id"`
	SongName    string    `json:"name"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
	ReleaseDate time.Time `json:"release_date"`
}
