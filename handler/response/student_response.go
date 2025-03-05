package response

import (
	"time"

	"github.com/google/uuid"
)

type StudentResponse struct {
	Id           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	LastName     string       `json:"last_name"`
	Registration string       `json:"registration"`
	User         UserResponse `json:"user"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
}
