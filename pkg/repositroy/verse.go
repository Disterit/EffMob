package repositroy

import (
	"EffMob/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VerseRepository struct {
	db *sqlx.DB
}

func NewVerseRepository(db *sqlx.DB) *VerseRepository {
	return &VerseRepository{db: db}
}

func (r *VerseRepository) GetVerses(songId int) (string, error) {
	var text string
	query := fmt.Sprintf(`SELECT text_song FROM %s WHERE id = $1`, SongTable)

	err := r.db.QueryRow(query, songId).Scan(&text)
	if err != nil {
		logger.Log.Error("error query row for get verses")
		return "", err
	}

	return text, nil
}
