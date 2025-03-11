package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase/validation"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type StudentUseCase struct {
	repo       repository.StudentRepository
	validators []validation.StudentRegisterValidator
}

func NewStudentUseCase(rp repository.StudentRepository, vl ...validation.StudentRegisterValidator) *StudentUseCase {
	return &StudentUseCase{
		repo:       rp,
		validators: vl,
	}
}

func (su *StudentUseCase) GetStudents() (*[]response.Student, error) {
	var st []response.Student
	err := su.repo.FindAll(&st)
	if err != nil {
		return nil, fmt.Errorf("error to get all students in database: %v", err)
	}
	return &st, nil
}

func (su *StudentUseCase) CreateStudent(std *request.Student) (string, error) {
	for _, v := range su.validators {
		if err := v.Validate(std); err != nil {
			return "", err
		}
	}
	student := std.New()
	err := su.repo.Save(student)
	if err != nil {
		return "", fmt.Errorf("error to create student in database: %v", err)
	}
	return student.ExternalId.String(), nil
}

func (su *StudentUseCase) GetStudentById(id string) (*response.Student, error) {
	var st response.Student
	err := su.repo.FindById(&st, id)
	if err != nil {
		return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
	}
	return &st, nil
}
