package repository

import (
	"database/sql"
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
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
	if err := row.Scan(&user.ExternalId, &user.Email, &user.Password); err != nil {
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

func (u *UserRepository) FindAll(stds *[]response.Student) error {
	queryFindAll := `select u.external_id, u.email, s.external_id, s.name, s.last_name, 
		s.registration, u.created_at, u.updated_at, s.created_at, s.updated_at from users u inner join students s on u.id = s.user_id;`
	rows, err := u.db.Query(queryFindAll)
	if err != nil {
		return fmt.Errorf("error to find all users: %v", err)
	}
	defer rows.Close()
	var student response.Student
	for rows.Next() {
		if err = rows.Scan(&student.User.ExternalId, &student.User.Email, &student.ExternalId, &student.Name, &student.LastName,
			&student.Registration, &student.User.CreatedAt, &student.User.UpdatedAt, &student.CreatedAt, &student.UpdatedAt); err != nil {
			return fmt.Errorf("error to scan data: %v", err)
		}
		*stds = append(*stds, student)
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
