package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type StudentUseCase struct {
	repo repository.StudentRepository
}

func NewStudentUseCase(rp repository.StudentRepository) *StudentUseCase {
	return &StudentUseCase{
		repo: rp,
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
	student := std.New()
	err := su.checkDuplication(std)
	if err != nil {
		return "", err
	}
	err = su.repo.Save(student)
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

func (su *StudentUseCase) checkDuplication(st *request.Student) error {
	var exists bool
	err := su.repo.ExistsByRegistration(&st.Registration, &exists)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("there is a student registered by this email")
	}
	return nil
}
