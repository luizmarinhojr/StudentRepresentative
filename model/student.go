package model

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Registration string    `json:"registration"`
	User         User      `json:"user"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func (st *Student) New() *Student {
	return &Student{}
}
