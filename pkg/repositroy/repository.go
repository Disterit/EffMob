package repositroy

import (
	"EffMob/models"
	"EffMob/pkg/handler"
	"github.com/jmoiron/sqlx"
)

type Group interface {
	CreateGroup(groupName string) (int, error)
	GetAllLibrary() (handler.OutputLibrary, error)
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
	Group
	Song
	Verse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Group: NewGroupRepository(db),
		Song:  NewSongRepository(db),
		Verse: NewVerseRepository(db),
	}
}
