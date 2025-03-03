package schema

import (
	"time"

	"github.com/google/uuid"
)

type Professor struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Registration string    `json:"registration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type ProfessorResponse struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Registration string    `json:"registration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
