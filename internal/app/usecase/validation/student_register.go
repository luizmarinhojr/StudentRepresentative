package validation

import (
	"fmt"
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
)

type StudentDuplicationByRegister struct {
	repo repository.StudentRepository
}

func NewStudentDuplicationByRegister(r repository.StudentRepository) *StudentDuplicationByRegister {
	return &StudentDuplicationByRegister{
		repo: r,
	}
}

func (st *StudentDuplicationByRegister) Validate(student *request.Student) error {
	var exists bool
	err := st.repo.ExistsByRegistration(&student.Registration, &exists)
	log.Println("STUDENT VALOR EXISTS:", exists)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("there is a student registered by this email")
	}
	return nil
}
