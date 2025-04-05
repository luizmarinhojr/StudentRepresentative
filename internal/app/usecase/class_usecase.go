package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type ClassUseCase struct {
	repo repository.ClassRepository
}

func NewClassUseCase(rep repository.ClassRepository) *ClassUseCase {
	return &ClassUseCase{
		repo: rep,
	}
}

func (cc *ClassUseCase) GetClassById(id string) (*response.Class, error) {
	var class response.Class
	err := cc.repo.FindById(&class, id)
	fmt.Println("CLASSE NO USECASE:", class.Name)
	if err != nil {
		return nil, fmt.Errorf("error to find by id: %v", err)
	}
	return &class, nil
}
