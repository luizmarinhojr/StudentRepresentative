package model

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	Id            int       `json:"id"`
	External_id   uuid.UUID `json:"external_id"`
	Name          string    `json:"name"`
	StartYear     int       `json:"start_year"`
	StartSemester int       `json:"start_semester"`
	EndYear       int       `json:"end_year"`
	EndSemester   int       `json:"end_semester"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func NewClass() *Class {
	return &Class{}
}
