package response

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ExternalId *uuid.UUID `json:"id"`
	Email      *string    `json:"email"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
