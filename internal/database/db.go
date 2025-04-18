package database

import (
	"database/sql"
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/config"
)

func OpenConnection() (*sql.DB, error) {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PSQL_HOST_DEV, config.PSQL_PORT_DEV, config.PSQL_USER_DEV, config.PSQL_PASS_DEV, config.PSQL_DBNAME_DEV)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic("Erro aqui")
	}
	err = db.Ping()

	return db, err
}
