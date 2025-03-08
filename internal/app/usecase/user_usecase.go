package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/auth"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/service"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
)

type UserUseCase struct {
	userRepo        repository.UserRepository
	studentRepo     repository.StudentRepository
	studentService  service.StudentService
	passwordService service.PasswordService
}

func NewUserUseCase(rpu repository.UserRepository, rps repository.StudentRepository, sv service.StudentService, ps service.PasswordService) *UserUseCase {
	return &UserUseCase{
		userRepo:        rpu,
		studentRepo:     rps,
		studentService:  sv,
		passwordService: ps,
	}
}

func (us *UserUseCase) SignIn(user *request.User) (*string, error) {
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

func (us *UserUseCase) SignUp(user *request.Login) error {
	var exists bool
	err := us.userRepo.ExistsByEmail(&user.Email, &exists)
	if err != nil {
		return fmt.Errorf("you do not have permission to sign up")
	}
	if exists {
		return fmt.Errorf("there is a student registered by this email")
	}
	err = us.studentService.IsStudentRegistered(&user.Registration)
	if err == nil {
		return fmt.Errorf("you do not have permission to sign up 2")
	}
	userModel := user.New()
	passwordHash, erro := us.passwordService.HashPassword(userModel.Password)
	if erro != nil {
		return fmt.Errorf("error to hash password: %v", erro)
	}
	userModel.Password = passwordHash
	err = us.userRepo.Save(userModel)
	if err != nil {
		return fmt.Errorf("error to save de user in database: %v", err)
	}
	err = us.studentRepo.UpdateUserByRegistration(userModel.Id, user.Registration)
	if err != nil {
		return err
	}
	return nil
}
