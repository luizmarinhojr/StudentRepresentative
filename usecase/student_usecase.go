package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/handler/request"
	"github.com/luizmarinhojr/StudentRepresentative/handler/response"
	"github.com/luizmarinhojr/StudentRepresentative/repository"
)

type StudentUseCase struct {
	repo repository.StudentRepository
}

func NewStudentUseCase(rp repository.StudentRepository) *StudentUseCase {
	return &StudentUseCase{
		repo: rp,
	}
}

func (su *StudentUseCase) GetStudents() (*[]response.StudentResponse, error) {
	var st []response.StudentResponse
	err := su.repo.SelectAllStudents(&st)
	if err != nil {
		return nil, fmt.Errorf("error to get all students in database: %v", err)
	}
	return &st, nil
}

func (su *StudentUseCase) CreateStudent(std *request.StudentRequest) (string, error) {
	student := std.New()
	err := su.checkDuplication(std)
	if err != nil {
		return "", err
	}
	err = su.repo.InsertIntoStudents(student)
	if err != nil {
		return "", fmt.Errorf("error to create student in database: %v", err)
	}
	return student.Id.String(), nil
}

func (su *StudentUseCase) GetStudentById(id string) (*response.StudentResponse, error) {
	var st response.StudentResponse
	err := su.repo.SelectStudentById(&st, id)
	if err != nil {
		return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
	}
	return &st, nil
}

func (su *StudentUseCase) checkDuplication(st *request.StudentRequest) error {
	registration := st.Registration
	var exists bool
	err := su.repo.ExistsStudentByRegistration(&registration, &exists)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("there is a student registered with that registration")
	}
	return nil
}
