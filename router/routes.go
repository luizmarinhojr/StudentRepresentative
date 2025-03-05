package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/handler"
	"github.com/luizmarinhojr/StudentRepresentative/repository"
	"github.com/luizmarinhojr/StudentRepresentative/usecase"
)

func InitializeRoutes(r *gin.Engine, db *sql.DB) {
	studentRepository := repository.NewStudentRepository(db)
	studentUseCase := usecase.NewStudentUseCase(*studentRepository)
	studentHandler := handler.NewStudentController(*studentUseCase)

	fromBase := r.Group("/api/v1")
	{
		{
			fromBase.POST("student", studentHandler.CreateStudent)
			fromBase.GET("/student/:id", studentHandler.GetStudentById)
			fromBase.GET("/students", studentHandler.GetStudents)
			fromBase.GET("/students/name/:name", studentHandler.GetAllStudentsByName)
			fromBase.GET("/students/registration/:registration", studentHandler.GetOneStudentByRegistration)
		}
	}

	r.Run()
}
