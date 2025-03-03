package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
)

func getConnection() *sql.DB {
	db, err := OpenConnection(false)
	if err != nil {
		panic("no connection to database")
	}
	return db
}

func GetAllStudents() (*[]schema.StudentResponse, error) {
	db := getConnection()
	var students []schema.StudentResponse
	rows, erro := db.Query("SELECT id, name, last_name, registration, created_at, updated_at FROM students;")
	if erro != nil {
		return nil, fmt.Errorf("the query isn't correct: %v", erro)
	}
	defer rows.Close()

	for rows.Next() {
		var st schema.StudentResponse
		if err := rows.Scan(&st.Id, &st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt); err != nil {
			return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
		}
		students = append(students, st)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("any error to finish: %v", err)
	}
	return &students, nil
}

func CreateStudent(student *schema.Student) (string, error) {
	db, err := getConnection().Begin()
	if err != nil {
		return "", fmt.Errorf("error to begin a transaction: %v", err)
	}
	var id string
	row := db.QueryRow("INSERT INTO students (name, last_name, registration) VALUES ($1, $2, $3) RETURNING id;", student.Name, student.LastName, student.Registration)
	if err := row.Scan(&id); err != nil {
		db.Rollback()
		return "", fmt.Errorf("it couldn't recovery the id correctly: %v", err)
	}
	db.Commit()
	return id, nil
}

func GetStudentById(id string) (*schema.StudentResponse, error) {
	db := getConnection()
	row := db.QueryRow("SELECT id, name, last_name, registration, created_at, updated_at FROM students WHERE id = $1", id)
	var st schema.StudentResponse
	if err := row.Scan(&st.Id, &st.Name, &st.LastName, &st.Registration, &st.CreatedAt, &st.UpdatedAt); err != nil {
		fmt.Println("erro aqui: ", err)
		return nil, fmt.Errorf("error to scan data: %v", err)
	}
	db.Close()
	return &st, nil
}
