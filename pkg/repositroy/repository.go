package repositroy

import (
	"EffMob/models"
	"github.com/jmoiron/sqlx"
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
