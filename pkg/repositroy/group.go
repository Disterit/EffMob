package repositroy

import (
	"EffMob/pkg/handler"
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

func (r *GroupRepository) GetAllLibrary() (handler.OutputLibrary, error) {

	return handler.OutputLibrary{}, nil
}
