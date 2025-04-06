package usecase

import (
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

func (cc *ClassUseCase) GetClasses() (*[]response.Class, error) {
	var classes []response.Class
	err := cc.repo.FindAll(&classes)
	if err != nil {
		return nil, err
	}
	return &classes, nil
}

func (cc *ClassUseCase) GetClassById(id string) (*response.Class, error) {
	var class response.Class
	err := cc.repo.FindById(&class, id)
	if err != nil {
		return nil, err
	}
	return &class, nil
}
