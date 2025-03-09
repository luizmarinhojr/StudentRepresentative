package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (st *Login) New() *model.User {
	return &model.User{
		Email:    st.Email,
		Password: []byte(st.Password),
	}
}

func (st *Login) Validate() error {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return err
	}
	return nil
}
