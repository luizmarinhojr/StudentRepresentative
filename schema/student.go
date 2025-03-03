package schema

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Registration string    `json:"registration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type StudentRequest struct {
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Registration string `json:"registration"`
}

type StudentResponse struct {
	Id           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	LastName     string     `json:"last_name"`
	Registration string     `json:"registration"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (st *StudentRequest) NewStudent() *Student {
	return &Student{
		Name:         st.Name,
		LastName:     st.LastName,
		Registration: st.Registration,
	}
}

// STUDENT'S METHODS
func (st *Student) QueryInsertInto() (string, []any) {
	return "INSERT INTO students (name, last_name, registration) VALUES ($1, $2, $3) RETURNING id", []any{st.Name, st.LastName, st.Registration}
}

func (st *Student) QuerySelectAll() string {
	return "SELECT id, name, last_name, registration, created_at, updated_at FROM students;"
}

func (st *Student) QuerySelectById() (string, []any) {
	return "SELECT id, name, last_name, registration, created_at, updated_at FROM students WHERE id = $1", []any{&st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt}
}

// STUDENTRESPONSE'S METHODS
func (st *StudentResponse) QueryInsertInto() (string, []any) {
	return "INSERT INTO students (name, last_name, registration) VALUES ($1, $2, $3) RETURNING id", []any{st.Name, st.LastName, st.Registration}
}

func (st *StudentResponse) QuerySelectAll() string {
	return "SELECT id, name, last_name, registration, created_at, updated_at FROM students;"
}

func (st *StudentResponse) QuerySelectById() (string, []any) {
	return "SELECT id, name, last_name, registration, created_at, updated_at FROM students WHERE id = $1", []any{&st.Id, &st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt}
}
