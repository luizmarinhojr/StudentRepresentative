package schema

type Repository interface {
	QuerySelectAll() string
	QuerySelectById() (string, []any)
	QueryInsertInto() (string, []any)
}

type Entities interface {
	StudentResponse | ProfessorResponse
}
