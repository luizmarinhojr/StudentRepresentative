package usecase

import (
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/repository"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
)

func GetStudents() (*[]schema.StudentResponse, error) {
	var st schema.StudentResponse
	students, err := repository.SelectAllDb(&schema.Student{}, &st)
	if err != nil {
		return nil, fmt.Errorf("error to get all students in database: %v", err)
	}
	return students, nil
}

func CreateStudent(std *schema.StudentRequest) (string, error) {
	student := std.NewStudent()
	id, err := repository.InsertIntoDb(student)
	if err != nil {
		return "", fmt.Errorf("error to create student in database: %v", err)
	}
	return id, nil
}

func GetStudentById(id string) (*schema.StudentResponse, error) {
	var st schema.StudentResponse
	err := repository.SelectById(&st, id)
	// student, err := repository.GetStudentById(id)
	if err != nil {
		return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
	}
	return &st, nil
}
