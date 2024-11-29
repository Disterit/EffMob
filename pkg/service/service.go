package service

import (
	"EffMob/models"
	"EffMob/pkg/repositroy"
)

type Group interface {
}

type Song interface {
	CreateSong(groupName, songName string) (int, error)
	GetAllSongs() ([]models.Song, error)
	GetSongById(id int) (models.Song, error)
	UpdateSong(id int, input models.UpdateSong) error
	DeleteSong(id int) error
}

type Verse interface {
	GetVerses(songId, verseId, limit int) (map[string]string, error)
}

type Service struct {
	Song
	Verse
}

func NewService(repo *repositroy.Repository) *Service {
	return &Service{
		Song:  NewSongService(repo),
		Verse: NewVerseService(repo),
	}
}
