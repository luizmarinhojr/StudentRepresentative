package handler

import "github.com/luizmarinhojr/StudentRepresentative/usecase"

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(us usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: us,
	}
}
