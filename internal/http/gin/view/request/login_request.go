package request

import "github.com/luizmarinhojr/StudentRepresentative/internal/app/model"

type Login struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	Registration string `json:"registration" validate:"required,min=12,max=12"`
}

func (st *Login) New() *model.User {
	return &model.User{
		Email:    st.Email,
		Password: []byte(st.Password),
	}
}
