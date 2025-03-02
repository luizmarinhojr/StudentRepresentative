package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/repository"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
)

func CreateStudent(std *schema.StudentRequest) (string, error) {
	student := std.NewStudent()
	id, err := repository.CreateStudent(student)
	if err != nil {
		return "", fmt.Errorf("error to create student in database: %v", err)
	}
	return id, nil
}

func GetStudentById(id string) (*schema.StudentResponse, error) {
	student, err := repository.GetStudentById(id)
	if err != nil {
		return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
	}
	return student, nil
}
