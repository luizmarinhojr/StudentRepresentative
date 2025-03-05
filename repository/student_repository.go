package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/handler/response"
	"github.com/luizmarinhojr/StudentRepresentative/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(con *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: con,
	}
}

func (sr *StudentRepository) SelectAllStudents(st *[]response.StudentResponse) error {
	querySelectAll := `SELECT s.id, s.name, s.last_name, s.registration, s.created_at, 
						s.updated_at, u.id, u.email, u.created_at, u.updated_at FROM students s left join users u on s.user_id = u.id;`
	var s response.StudentResponse
	rows, err := sr.db.Query(querySelectAll)
	if err != nil {
		return fmt.Errorf("error to select all: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&s.Id, &s.Name, &s.LastName, &s.Registration, &s.CreatedAt, &s.UpdatedAt,
			&s.User.Id, &s.User.Email, &s.User.CreatedAt, &s.User.UpdatedAt); err != nil {
			return fmt.Errorf("error to catch data: %v", err)
		}
		*st = append(*st, s)
	}
	return nil
}

func (sr *StudentRepository) SelectStudentById(st *response.StudentResponse, id string) error {
	querySelectById := `SELECT s.id, s.name, s.last_name, s.registration, s.created_at, 
						s.updated_at, u.id, u.email, u.created_at, u.updated_at FROM students s full join users u on s.user_id = u.id where s.id = $1`
	row := sr.db.QueryRow(querySelectById, &id)
	if err := row.Scan(&st.Id, &st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt,
		&st.User.Id, &st.User.Email, &st.User.CreatedAt, &st.User.UpdatedAt); err != nil {
		return fmt.Errorf("error to catch student by id: %v", err)
	}
	return nil
}

func (sr *StudentRepository) InsertIntoStudents(st *model.Student) error {
	queryInsertInto := "INSERT INTO students (name, last_name, registration) VALUES ($1, $2, $3) RETURNING id"
	row := sr.db.QueryRow(queryInsertInto, &st.Name, &st.LastName, &st.Registration)
	if err := row.Scan(&st.Id); err != nil {
		return fmt.Errorf("error to insert student: %v", err)
	}
	return nil
}

func (sr *StudentRepository) ExistsStudentByRegistration(registration *string, exists *bool) error {
	queryExistsByRegistration := "SELECT EXISTS (SELECT 1 FROM students WHERE registration = $1);"
	row := sr.db.QueryRow(queryExistsByRegistration, registration)
	if err := row.Scan(exists); err != nil {
		return err
	}
	return nil
}
