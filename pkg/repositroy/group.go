package repositroy

import (
	"EffMob/logger"
	"EffMob/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GroupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(groupName string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (group_name) VALUES ($1) RETURNING id", GroupTable)
	row := r.db.QueryRow(query, groupName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *GroupRepository) GetAllLibrary() (map[string][]models.Song, error) {

	outputLibraries := make(map[string][]models.Song)

	query := `SELECT
    			groups.group_name, 
    			songs.id,
    			groups.id,
				songs.song_name, 
				songs.text_song, 
				songs.link, 
				songs.release_date 
			FROM groups
			JOIN songs ON groups.id=songs.group_id
			ORDER BY groups.group_name; `

	rows, err := r.db.Query(query)
	if err != nil {
		logger.Log.Error("error to get library from database")
		return nil, err
	}

	for rows.Next() {
		var group string
		var song models.Song
		err = rows.Scan(
			&group,
			&song.Id,
			&song.GroupId,
			&song.SongName,
			&song.Text,
			&song.Link,
			&song.ReleaseDate,
		)

		if err != nil {
			logger.Log.Error("error to get library from database")
			return nil, err
		}

		outputLibraries[group] = append(outputLibraries[group], song)
	}

	return outputLibraries, nil
}

func (r *GroupRepository) GetAllSongGroupById(id int) (map[string][]models.Song, error) {
	outputGroupSongs := make(map[string][]models.Song)

	query := `SELECT
    			groups.group_name, 
    			songs.id,
    			groups.id,
				songs.song_name, 
				songs.text_song, 
				songs.link, 
				songs.release_date 
			FROM groups
			JOIN songs ON groups.id=songs.group_id
			WHERE songs.group_id = $1;  `

	rows, err := r.db.Query(query, id)
	if err != nil {
		logger.Log.Error("error to get song group from database")
		return nil, err
	}
	for rows.Next() {
		var group string
		var song models.Song
		err = rows.Scan(
			&group,
			&song.Id,
			&song.GroupId,
			&song.SongName,
			&song.Text,
			&song.Link,
			&song.ReleaseDate,
		)
		if err != nil {
			logger.Log.Error("error to get song group from database")
			return nil, err
		}
		outputGroupSongs[group] = append(outputGroupSongs[group], song)
	}

	return outputGroupSongs, nil
}

func (r *GroupRepository) UpdateGroup(id int, input models.Group) error {

	query := fmt.Sprintf("UPDATE %s SET group_name = $1 WHERE id = $2", GroupTable)
	_, err := r.db.Exec(query, input.GroupName, id)
	if err != nil {
		logger.Log.Error("error to update group from database", err.Error())
		return err
	}

	return nil
}

func (r *GroupRepository) DeleteGroup(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", GroupTable)

	_, err := r.db.Exec(query, id)
	if err != nil {
		logger.Log.Error("error to delete group from database", err.Error())
		return err
	}

	return nil
}
