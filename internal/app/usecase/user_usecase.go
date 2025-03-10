package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/auth"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/service"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase/validator"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type UserUseCase struct {
	userRepo          repository.UserRepository
	studentRepo       repository.StudentRepository
	passwordService   service.PasswordService
	studentValidators []validator.StudentRegisterValidator
	userValidators    []validator.UserRegisterValidator
}

func NewUserUseCase(rpu repository.UserRepository, rps repository.StudentRepository,
	ps service.PasswordService, sv []validator.StudentRegisterValidator, vl ...validator.UserRegisterValidator) *UserUseCase {
	return &UserUseCase{
		userRepo:          rpu,
		studentRepo:       rps,
		passwordService:   ps,
		userValidators:    vl,
		studentValidators: sv,
	}
}

func (us *UserUseCase) SignIn(user *request.Login) (*string, error) {
	userDb := user.New()
	err := us.userRepo.FindByEmail(userDb)
	if err != nil {
		return nil, err
	}
	err = us.passwordService.CheckPasswordHash([]byte(user.Password), userDb.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %v", err.Error())
	}
	token, erro := auth.GenerateJWT(userDb)
	if erro != nil {
		return nil, fmt.Errorf("error to generate the jwt token: %v", erro)
	}
	return &token, nil
}

func (us *UserUseCase) SignUp(user *request.User) (string, error) {
	for _, v := range us.userValidators {
		if err := v.Validate(user); err != nil {
			return "", err
		}
	}
	userModel := user.New()
	passwordHash, erro := us.passwordService.HashPassword(userModel.Password)
	if erro != nil {
		return "", fmt.Errorf("error to hash password: %v", erro)
	}
	userModel.Password = passwordHash
	if err := us.userRepo.Save(userModel); err != nil {
		return "", err
	}
	if err := us.studentRepo.UpdateUserByRegistration(userModel.Id, user.Registration); err != nil {
		return "", err
	}

	return userModel.ExternalId.String(), nil
}

func (us *UserUseCase) GetUsers() (*[]response.Student, error) {
	var students []response.Student
	err := us.userRepo.FindAll(&students)
	if err != nil {
		return nil, err
	}
	return &students, nil
}
