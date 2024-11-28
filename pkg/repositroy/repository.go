package repositroy

import "github.com/jmoiron/sqlx"

type Song interface {
	CreateSong(groupName, songName string) (int, error)
}

type Verse interface {
}

type Repository struct {
	Song
	Verse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song: NewSongRepository(db),
	}
}
