package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/service"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/handler"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/middleware"
)

func InitializeRoutes(r *gin.Engine, db *sql.DB) {
	studentRepository := repository.NewStudentRepository(db)
	studentUseCase := usecase.NewStudentUseCase(*studentRepository)
	studentHandler := handler.NewStudentController(*studentUseCase)

	studentService := service.NewStudentService(*studentRepository)
	passwordService := service.NewPasswordService()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(*userRepository, *studentRepository, *studentService, *passwordService)
	userHandler := handler.NewUserHandler(*userUseCase)

	fromBase := r.Group("/api/v1")
	{
		{
			fromBase.POST("student", studentHandler.CreateStudent)
			fromBase.GET("/student/:id", studentHandler.GetStudentById)
			fromBase.GET("/students", middleware.CheckAuth, studentHandler.GetStudents)
			fromBase.GET("/students/name/:name", studentHandler.GetAllStudentsByName)
			fromBase.GET("/students/registration/:registration", studentHandler.GetOneStudentByRegistration)
		}
		{
			fromBase.POST("signin", userHandler.SignIn)
			fromBase.POST("signup", userHandler.SignUp)
		}
	}

	r.Run()
}
