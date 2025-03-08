package service

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
)

type StudentService struct {
	repo repository.StudentRepository
}

func NewStudentService(rp repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: rp,
	}
}

func (s *StudentService) IsStudentRegistered(registration *string) error {
	var exists bool
	err := s.repo.ExistsByRegistration(registration, &exists)
	if exists {
		return fmt.Errorf("there is no student by this registration: %v", err)
	}
	return nil
}
