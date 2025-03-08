package response

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ExternalId   uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	LastName     string     `json:"last_name"`
	Registration string     `json:"registration"`
	User         User       `json:"user"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
