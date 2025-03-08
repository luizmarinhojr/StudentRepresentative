package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (us *User) New() *model.User {
	return &model.User{
		Email:    us.Email,
		Password: []byte(us.Password),
	}
}

func (us *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(us)
	if err != nil {
		return err
	}
	return nil
}
