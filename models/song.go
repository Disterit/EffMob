package models

import (
	"database/sql"
	"errors"
)

type Song struct {
	Id          int            `json:"id" db:"id"`
	GroupId     int            `json:"group_id" db:"group_id"`
	SongName    sql.NullString `json:"name" db:"song_name"`
	Text        sql.NullString `json:"text" db:"text_song"`
	Link        sql.NullString `json:"link" db:"link"`
	ReleaseDate sql.NullTime   `json:"release_date" db:"release_date"`
}

type UpdateSong struct {
	SongName    *string `json:"name" db:"name"`
	Text        *string `json:"text" db:"text_song"`
	Link        *string `json:"link" db:"link"`
	ReleaseDate *string `json:"release_date" db:"release_date"`
}

func (s *UpdateSong) Validate() error {
	if s.SongName == nil && s.Text == nil && s.Link == nil && s.ReleaseDate == nil {
		return errors.New("either title or description must be provided")
	}

	return nil
}
