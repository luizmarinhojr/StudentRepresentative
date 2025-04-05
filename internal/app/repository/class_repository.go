package repository

import (
	"database/sql"
	"fmt"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/model"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/response"
)

type ClassRepository struct {
	db *sql.DB
}

func NewClassRepository(con *sql.DB) *ClassRepository {
	return &ClassRepository{
		db: con,
	}
}

func (cl *ClassRepository) FindAll(cr *[]response.Class) error {
	queryFindAll := `SELECT external_id, name, start_year, start_semester, end_year, end_semester, created_at, updated_at FROM classes;`
	rows, err := cl.db.Query(queryFindAll)
	if err != nil {
		return fmt.Errorf("error to select all: %v", err)
	}
	defer rows.Close()
	var class response.Class
	for rows.Next() {
		if err := rows.Scan(&class.External_id, &class.Name, &class.StartYear, &class.StartSemester, &class.EndYear, &class.EndSemester,
			&class.CreatedAt, &class.UpdatedAt); err != nil {
			return fmt.Errorf("error to catch data: %v", err)
		}
		*cr = append(*cr, class)
	}
	return nil
}

func (cl *ClassRepository) FindById(cr *response.Class, id string) error {
	queryFindById := `SELECT external_id, name, start_year, start_semester, end_year, end_semester, created_at, updated_at FROM classes
		WHERE external_id = $1`
	row := cl.db.QueryRow(queryFindById, &id)
	if err := row.Scan(&cr.External_id, &cr.Name, &cr.StartYear, &cr.StartSemester, &cr.EndYear, &cr.EndSemester,
		&cr.CreatedAt, &cr.UpdatedAt); err != nil {
		return fmt.Errorf("error to catch class by id: %v", err)
	}
	return nil
}

func (cl *ClassRepository) Save(cr *model.Class) error {
	queryInsertInto := `INSERT INTO classes(name, start_year, start_semester, end_year, end_semester)
						 VALUES($1, $2, $3, $4, $5) RETURNING external_id`
	transaction, err := cl.db.Begin()
	row := transaction.QueryRow(queryInsertInto, &cr.Name, &cr.StartYear, &cr.StartSemester, &cr.EndYear, &cr.EndSemester)
	if err := row.Scan(&cr.External_id); err != nil {
		if transaction.Rollback() != nil {
			return fmt.Errorf("error to rollback the changes: %v", err)
		}
		return fmt.Errorf("error to catch id: %v", err)
	}
	if transaction.Commit() != nil {
		return fmt.Errorf("error to commit the changes: %v", err)
	}
	return nil
}
