package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/model"
)

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (us *UserRequest) New() *model.User {
	return &model.User{
		Email:    us.Email,
		Password: us.Password,
	}
}

func (us *UserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(us)
	if err != nil {
		return err
	}
	return nil
}
