package repositroy

import (
	"EffMob/logger"
	"EffMob/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type SongRepository struct {
	db *sqlx.DB
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) CreateSong(groupName, songName string, songInfo *models.SongInfo) (int, error) {
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

	query = fmt.Sprintf(`INSERT INTO %s (group_id, song_name, text_song, link, release_date) 
								VALUES ($1, $2, $3, $4, TO_DATE($5, 'DD.MM.YYYY')) RETURNING id`, SongTable)

	res = tx.QueryRow(query, id, songName, songInfo.Text, songInfo.Link, songInfo.ReleaseDate)
	if err = res.Scan(&id); err != nil {
		tx.Rollback()
		logger.Log.Error("error to insert song or conflict resolution", err.Error())
		return 0, err
	}

	tx.Commit()

	return id, nil
}

func (r *SongRepository) GetAllSongs() ([]models.Song, error) {
	var songs []models.Song

	query := fmt.Sprintf(`SELECT * FROM %s`, SongTable)
	rows, err := r.db.Queryx(query)
	if err != nil {
		logger.Log.Error("error to query all songs", err.Error())
		return nil, err
	}

	for rows.Next() {
		var song models.Song
		if err := rows.StructScan(&song); err != nil {
			logger.Log.Error("error to scan all song", err.Error())
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) GetSongById(id int) (models.Song, error) {
	var song models.Song

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, SongTable)
	res := r.db.QueryRowx(query, id)
	if err := res.StructScan(&song); err != nil {
		logger.Log.Error("error to scan song", err.Error())
		return song, err
	}

	return song, nil
}

func (r *SongRepository) UpdateSong(id int, input models.UpdateSong) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.SongName != nil {
		setValues = append(setValues, fmt.Sprintf("song_name=$%d", argId))
		args = append(args, *input.SongName)
		argId++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text_song=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	if input.Link != nil {
		setValues = append(setValues, fmt.Sprintf("link=$%d", argId))
		args = append(args, *input.Link)
		argId++
	}

	if input.ReleaseDate != nil {
		parsedDate, err := time.Parse("01-02-2006", *input.ReleaseDate)
		if err != nil {
			return fmt.Errorf("invalid release_date format: %w", err)
		}
		setValues = append(setValues, fmt.Sprintf("release_date=$%d", argId))
		args = append(args, parsedDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", SongTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update song from repository: %w", err)
	}

	return nil
}

func (r *SongRepository) DeleteSong(id int) error {

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, SongTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete song from repository: %w", err)
	}

	return nil
}
