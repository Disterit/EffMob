package repositroy

import (
	"EffMob/logger"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	SongTable  = "songs"
	GroupTable = "groups"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func Connection(config Config) *sqlx.DB {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.Database, config.SSLMode))
	if err != nil {
		logger.Log.Error("error to connect db", err.Error())
		return nil
	}

	return db
}
