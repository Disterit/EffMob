package repositroy

import (
	"EffMob/models"
	"github.com/jmoiron/sqlx"
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
	GetVerses(songId int) (string, error)
}

type Repository struct {
	Song
	Verse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song:  NewSongRepository(db),
		Verse: NewVerseRepository(db),
	}
}
