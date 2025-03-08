package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         int64     `json:"id"`
	ExternalId uuid.UUID `json:"external_id"`
	Email      string    `json:"email"`
	Password   []byte    `json:"password"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func NewUser() *User {
	return &User{}
}
