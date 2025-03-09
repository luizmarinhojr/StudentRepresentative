package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/dependencies"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/middleware"
)

func InitializeRoutes(r *gin.Engine, dependency dependencies.Dependencies) {

	fromBase := r.Group("/api/v1")
	{
		{
			fromBase.POST("student", dependency.StudentHandler.CreateStudent)
			fromBase.GET("/student/:id", dependency.StudentHandler.GetStudentById)
			fromBase.GET("/students", middleware.CheckAuth, dependency.StudentHandler.GetStudents)
			fromBase.GET("/students/name/:name", dependency.StudentHandler.GetAllStudentsByName)
			fromBase.GET("/students/registration/:registration", dependency.StudentHandler.GetOneStudentByRegistration)
		}
		{
			fromBase.POST("signin", dependency.UserHandler.SignIn)
			fromBase.POST("signup", dependency.UserHandler.SignUp)
		}
	}

	r.Run()
}
