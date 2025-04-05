package validation

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
)

type UserDuplicationByEmail struct {
	repo repository.UserRepository
}

func NewUserDuplicationByEmail(r repository.UserRepository) *UserDuplicationByEmail {
	return &UserDuplicationByEmail{
		repo: r,
	}
}

func (ud *UserDuplicationByEmail) Validate(user *request.User) error {
	var exists bool
	err := ud.repo.ExistsByEmail(&user.Email, &exists)
	if err != nil {
		return fmt.Errorf("error to check existence in database")
	}
	if exists {
		return fmt.Errorf("there is a student registered by this email")
	}
	return nil
}

type UserIsStudentExists struct {
	repo repository.StudentRepository
}

func NewUserIsStudentExists(r repository.StudentRepository) *UserIsStudentExists {
	return &UserIsStudentExists{
		repo: r,
	}
}

func (us *UserIsStudentExists) Validate(user *request.User) error {
	var exists bool
	if err := us.repo.ExistsByRegistration(&user.Registration, &exists); err != nil {
		return fmt.Errorf("error to check student's existence in database")
	}
	if !exists {
		return fmt.Errorf("there is no students registered by this email")
	}
	return nil
}

type StudentHaveUser struct {
	repo repository.StudentRepository
}

func NewStudentHaveUser(r repository.StudentRepository) *StudentHaveUser {
	return &StudentHaveUser{
		repo: r,
	}
}

func (us *StudentHaveUser) Validate(user *request.User) error {
	var exists bool
	if err := us.repo.ExistsUserByRegistration(&user.Registration, &exists); err != nil {
		return fmt.Errorf("error to check user student in database")
	}
	if exists {
		return fmt.Errorf("there is a user related for this registration")
	}
	return nil
}

type UserValidationByLastName struct {
	repo repository.StudentRepository
}

func NewUserValidationByLastName(r repository.StudentRepository) *UserValidationByLastName {
	return &UserValidationByLastName{
		repo: r,
	}
}

func (uv *UserValidationByLastName) Validate(user *request.User) error {
	var exists bool
	if err := uv.repo.ExistsStudentByLastNameAndRegistration(&user.LastName, &user.Registration, &exists); err != nil {
		return fmt.Errorf("error to query the database")
	}
	if !exists {
		return fmt.Errorf("there is no students registered by this user and last name")
	}
	return nil
}
