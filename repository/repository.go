package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/config"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
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

func GetConnection() *sql.DB {
	db, err := OpenConnection(false)
	if err != nil {
		panic("no connection to database")
	}
	return db
}

func InsertIntoDb(entity schema.Repository) (string, error) {
	db := GetConnection()
	transaction, err := db.Begin()
	if err != nil {
		return "", err
	}
	query, values := entity.QueryInsertInto()
	row := transaction.QueryRow(query, values...)
	var id string
	if erro := row.Scan(&id); erro != nil {
		transaction.Rollback()
		return "", fmt.Errorf("error to insert in database: %v", erro)
	}
	transaction.Commit()
	db.Close()
	return id, nil
}

func SelectAllDb[T schema.Entities](entity schema.Repository, response *T) ([]T, error) {
	db := getConnection()
	sqlQuery := entity.QuerySelectAll()
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error to select all: %v", err)
	}
	pointers := schema.GetFieldPointers(response)
	var list []T
	for rows.Next() {
		if erro := rows.Scan(pointers...); erro != nil {
			return nil, fmt.Errorf("error to select all: %v", erro)
		}
		list = append(list, *response)
	}
	rows.Close()
	db.Close()
	return list, nil
}

func SelectById(entity schema.Repository, id string) error {
	db := GetConnection()
	query, fields := entity.QuerySelectById()
	row := db.QueryRow(query, id)
	if err := row.Scan(fields...); err != nil {
		return fmt.Errorf("error to select all: %v", err)
	}
	db.Close()
	return nil
}
