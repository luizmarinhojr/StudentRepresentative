package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(con *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: con,
	}
}

func (sr *StudentRepository) FindAll(st *[]response.Student) error {
	querySelectAll := `SELECT s.external_id, s.name, s.last_name, s.registration, s.created_at, 
						s.updated_at, u.external_id, u.email, u.created_at, u.updated_at 
						FROM students s left join users u on s.user_id = u.id;`
	var s response.Student
	rows, err := sr.db.Query(querySelectAll)
	if err != nil {
		return fmt.Errorf("error to select all: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&s.ExternalId, &s.Name, &s.LastName, &s.Registration, &s.CreatedAt, &s.UpdatedAt,
			&s.User.ExternalId, &s.User.Email, &s.User.CreatedAt, &s.User.UpdatedAt); err != nil {
			return fmt.Errorf("error to catch data: %v", err)
		}
		*st = append(*st, s)
	}
	return nil
}

func (sr *StudentRepository) FindById(st *response.Student, id string) error {
	querySelectById := `SELECT s.external_id, s.name, s.last_name, s.registration, s.created_at, 
						s.updated_at, u.external_id, u.email, u.created_at, u.updated_at FROM students s 
						full join users u on s.user_id = u.id where s.external_id = $1`
	row := sr.db.QueryRow(querySelectById, &id)
	if err := row.Scan(&st.ExternalId, &st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt,
		&st.User.ExternalId, &st.User.Email, &st.User.CreatedAt, &st.User.UpdatedAt); err != nil {
		return fmt.Errorf("error to catch student by id: %v", err)
	}
	return nil
}

func (sr *StudentRepository) Save(st *model.Student) error {
	queryInsertInto := "INSERT INTO students (name, last_name, registration) VALUES ($1, $2, $3) RETURNING external_id"
	transaction, err := sr.db.Begin()
	if err != nil {
		return fmt.Errorf("error to start transaction: %v", err)
	}
	row := transaction.QueryRow(queryInsertInto, &st.Name, &st.LastName, &st.Registration)
	if err := row.Scan(&st.ExternalId); err != nil {
		transaction.Rollback()
		return fmt.Errorf("error to insert student: %v", err)
	}
	err = transaction.Commit()
	if err != nil {
		return fmt.Errorf("error to commit the changes: %v", err)
	}
	return nil
}

func (sr *StudentRepository) ExistsByRegistration(registration *string, exists *bool) error {
	queryExistsByRegistration := "SELECT EXISTS (SELECT 1 FROM students WHERE registration = $1);"
	row := sr.db.QueryRow(queryExistsByRegistration, *registration)
	err := row.Scan(exists)
	if err != nil {
		return err
	}
	return nil
}

func (sr *StudentRepository) UpdateUserByRegistration(userId int64, registration string) error {
	queryUpdateUserByRegistration := `UPDATE students SET user_id = $1 WHERE registration = $2`
	_, err := sr.db.Exec(queryUpdateUserByRegistration, userId, registration)
	if err != nil {
		return fmt.Errorf("error to update students table: %v", err)
	}
	return nil
}
