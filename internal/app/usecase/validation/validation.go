package validation

import (
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
)

type StudentRegisterValidator interface {
	// Return error if it is valid. Return nil if it is not valid
	Validate(student *request.Student) error
}

type UserRegisterValidator interface {
	// Return error if it is valid. Return nil if it is not valid
	Validate(user *request.User) error
}
