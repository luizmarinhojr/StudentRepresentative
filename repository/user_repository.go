package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}
