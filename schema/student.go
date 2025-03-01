package schema

import (
	"github.com/google/uuid"
)

type Student struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	Registration int64     `json:"registration"`
	// Subject      []subject.Subject
}

func (st *Student) NewStudent(name, lastName string, registration int64) *Student {
	return &Student{
		Id:           uuid.New(),
		Name:         name,
		LastName:     lastName,
		Registration: registration,
		// Subject:      subjects,
	}
}

type StudentRequest struct {
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Registration int64  `json:"registration"`
}

// func (st *Student) AddNewSubject(subject subject.Subject) {
// 	st.Subject = append(st.Subject, subject)
// }
