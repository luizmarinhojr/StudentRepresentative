package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
)

func getConnection() *sql.DB {
	db, err := OpenConnection()
	if err != nil {
		panic("no connection to database")
	}
	return db
}

func GetAllStudents() ([]schema.Student, error) {
	db := getConnection()
	var students []schema.Student
	rows, erro := db.Query("SELECT id, name, last_name, registration FROM students")
	if erro != nil {
		return nil, fmt.Errorf("the query isn't correct: %v", erro)
	}
	defer rows.Close()

	for rows.Next() {
		var st schema.Student
		if err := rows.Scan(&st.Id, &st.Name, &st.LastName, &st.Registration); err != nil {
			return nil, fmt.Errorf("it couldn't recovery the data correctly: %v", err)
		}
		students = append(students, st)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("any error to finish: %v", err)
	}
	return students, nil
}

func CreateStudent() {
	// db := getConnection()

}
