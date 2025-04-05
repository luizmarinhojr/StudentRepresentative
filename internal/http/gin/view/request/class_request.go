package request

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
)

type Class struct {
	Name          string `json:"name" validate:"required,min=8"`
	StartYear     int    `json:"start_year" validate:"required"`
	StartSemester int    `json:"start_semester" validate:"required,min=1,max=2"`
	EndYear       int    `json:"end_year" validate:"required,min=8"`
	EndSemester   int    `json:"end_semester" validate:"required,min=1,max=2"`
}

func (cl *Class) New() *model.Class {
	return &model.Class{
		Name:          cl.Name,
		StartYear:     cl.StartYear,
		StartSemester: cl.StartSemester,
		EndYear:       cl.EndYear,
		EndSemester:   cl.EndSemester,
	}
}

func (cl *Class) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("greater-than", cl.validateStartYear)
	validate.RegisterValidation("example", cl.validateEndYear)
	validate.RegisterStructValidation(cl.validateCombinedSemesterYear, cl)
	err := validate.Struct(cl)
	if err != nil {
		return err
	}
	return nil
}

func (cl *Class) validateStartYear(fl validator.FieldLevel) bool {
	year := fl.Field().Int()
	yearNow := int64(time.Now().Year())
	return year >= yearNow-5 && year <= yearNow+5
}

func (cl *Class) validateEndYear(fl validator.FieldLevel) bool {
	year := fl.Field().Int()
	yearNow := int64(time.Now().Year())
	return year >= int64(cl.StartYear) && year <= yearNow+7
}

func (cl *Class) validateCombinedSemesterYear(clt validator.StructLevel) {
	if (cl.EndYear < cl.StartYear) || (cl.EndYear == cl.StartYear && cl.EndSemester < cl.StartSemester) {
		clt.ReportError(cl.EndYear, "EndYear", "endyear", "combinedSemesterYear", "")
	}
}
