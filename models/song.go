package models

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type SongInfo struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

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

func (s Song) MarshalJSON() ([]byte, error) {
	type Alias Song
	return json.Marshal(&struct {
		Id          int     `json:"id"`
		GroupId     int     `json:"group_id"`
		SongName    *string `json:"name"`
		Text        *string `json:"text"`
		Link        *string `json:"link"`
		ReleaseDate *string `json:"release_date"`
	}{
		Id:          s.Id,
		GroupId:     s.GroupId,
		SongName:    nullStringToPtr(s.SongName),
		Text:        nullStringToPtr(s.Text),
		Link:        nullStringToPtr(s.Link),
		ReleaseDate: nullTimeToPtr(s.ReleaseDate),
	})
}

func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullTimeToPtr(nt sql.NullTime) *string {
	if nt.Valid {
		t := nt.Time.Format("2006-01-02T15:04:05Z")
		return &t
	}
	return nil
}
