package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/config"
)

func TestConnection() error {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PSQL_HOST_DEV, config.PSQL_PORT_DEV, config.PSQL_USER_DEV, config.PSQL_PASS_DEV, config.PSQL_DBNAME_DEV)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic("Erro aqui")
	}
	err = db.Ping()
	db.Close()
	return err
}

func OpenConnection(isTest bool) (*sql.DB, error) {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PSQL_HOST_DEV, config.PSQL_PORT_DEV, config.PSQL_USER_DEV, config.PSQL_PASS_DEV, config.PSQL_DBNAME_DEV)
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic("Erro aqui")
	}
	err = db.Ping()
	if isTest {
		db.Close()
	}
	return db, err
}
