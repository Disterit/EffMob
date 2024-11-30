package service

import (
	"EffMob/models"
	"EffMob/pkg/repositroy"
)

type Group interface {
	CreateGroup(groupName string) (int, error)
	GetAllLibrary() (map[string][]models.Song, error)
	GetAllSongGroupById(id int) (map[string][]models.Song, error)
	UpdateGroup(id int, input models.Group) error
	DeleteGroup(id int) error
}

type Song interface {
	CreateSong(groupName, songName string, songInfo *models.SongInfo) (int, error)
	GetAllSongs() ([]models.Song, error)
	GetSongById(id int) (models.Song, error)
	UpdateSong(id int, input models.UpdateSong) error
	DeleteSong(id int) error
}

type Verse interface {
	GetVerses(songId, verseId, limit int) (map[string]string, error)
}

type Service struct {
	Group
	Song
	Verse
}

func NewService(repo *repositroy.Repository) *Service {
	return &Service{
		Group: NewGroupService(repo),
		Song:  NewSongService(repo),
		Verse: NewVerseService(repo),
	}
}
