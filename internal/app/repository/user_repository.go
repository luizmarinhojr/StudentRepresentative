package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (u *UserRepository) FindByEmail(user *model.User) error {
	query := `SELECT external_id, email, pass FROM users WHERE email = $1;`
	row := u.db.QueryRow(query, user.Email)
	log.Println("EMAIL DO USER:", user.Email)
	if err := row.Scan(&user.ExternalId, &user.Email, &user.Password); err != nil {
		log.Println("Entrou nesse erro")
		return err
	}
	return nil
}

func (u *UserRepository) ExistsByEmail(email *string, exists *bool) error {
	queryExistsByEmail := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1);"
	row := u.db.QueryRow(queryExistsByEmail, *email)
	if err := row.Scan(exists); err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Save(user *model.User) error {
	queryInsertInto := `INSERT INTO users(email, pass) VALUES($1, $2) RETURNING external_id, id;`
	transaction, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("error to start transaction: %v", err)
	}
	row := transaction.QueryRow(queryInsertInto, user.Email, user.Password)
	if err := row.Scan(&user.ExternalId, &user.Id); err != nil {
		transaction.Rollback()
		return fmt.Errorf("error to save the user in database: %v", err)
	}
	err = transaction.Commit()
	if err != nil {
		return fmt.Errorf("error to commit the changes: %v", err)
	}
	return nil
}
