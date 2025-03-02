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
