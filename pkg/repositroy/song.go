package repositroy

import (
	"EffMob/logger"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SongRepository struct {
	db *sqlx.DB
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) CreateSong(groupName, songName string) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		logger.Log.Error("error to begin transaction")
		return 0, err
	}

	query := fmt.Sprintf(`SELECT id FROM %s WHERE group_name = $1`, GroupTable)
	res := tx.QueryRow(query, groupName)
	if err := res.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			query = fmt.Sprintf(`INSERT INTO %s (group_name) VALUES ($1) RETURNING id`, GroupTable)
			res = tx.QueryRow(query, groupName)
			if err := res.Scan(&id); err != nil {
				tx.Rollback()
				logger.Log.Error("error to insert new group or getting group id", err.Error())
				return 0, err
			}
		} else {
			tx.Rollback()
			logger.Log.Error("error to query group id", err.Error())
			return 0, err
		}
	}

	query = fmt.Sprintf(`INSERT INTO %s (group_id, song_name) VALUES ($1, $2) ON CONFLICT (group_id, song_name) DO NOTHING RETURNING id`, SongTable)
	res = tx.QueryRow(query, id, songName)
	if err = res.Scan(&id); err != nil {
		tx.Rollback()
		logger.Log.Error("error to insert song or conflict resolution", err.Error())
		return 0, err
	}

	tx.Commit()

	return id, nil
}
