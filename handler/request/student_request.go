package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/model"
)

type StudentRequest struct {
	Name         string `json:"name" validate:"required,min=2,max=60"`
	LastName     string `json:"last_name" validate:"required,min=5,max=60"`
	Registration string `json:"registration" validate:"required,min=12,max=12"`
}

func (st *StudentRequest) New() *model.Student {
	return &model.Student{
		Name:         st.Name,
		LastName:     st.LastName,
		Registration: st.Registration,
	}
}

func (st *StudentRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return err
	}
	return nil
}
