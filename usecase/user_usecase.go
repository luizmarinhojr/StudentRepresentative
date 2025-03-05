package usecase

import "github.com/luizmarinhojr/StudentRepresentative/repository"

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUserCase(rp repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: rp,
	}
}
