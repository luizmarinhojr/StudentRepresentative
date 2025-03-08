package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
)

type Student struct {
	Name         string `json:"name" validate:"required,min=2,max=50"`
	LastName     string `json:"last_name" validate:"required,min=5,max=50"`
	Registration string `json:"registration" validate:"required,min=12,max=12"`
}

func (st *Student) New() *model.Student {
	return &model.Student{
		Name:         st.Name,
		LastName:     st.LastName,
		Registration: st.Registration,
	}
}

func (st *Student) Validate() error {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return err
	}
	return nil
}
