package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/dependencies"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/middleware"
)

func InitializeRoutes(r *gin.Engine, dependency dependencies.Dependencies) {

	baseURL := r.Group("/api/v1")
	{
		{
			baseURL.POST("student", dependency.StudentHandler.CreateStudent)
			baseURL.GET("/student/:id", dependency.StudentHandler.GetStudentById)
			baseURL.GET("/students", middleware.CheckAuth, dependency.StudentHandler.GetStudents)
			baseURL.GET("/students/name/:name", dependency.StudentHandler.GetAllStudentsByName)
			baseURL.GET("/students/registration/:registration", dependency.StudentHandler.GetOneStudentByRegistration)
		}
		{
			baseURL.POST("signin", dependency.UserHandler.SignIn)
			baseURL.POST("signup", dependency.UserHandler.SignUp)
			baseURL.GET("users", dependency.UserHandler.GetUsers)
		}
		{
			baseURL.GET("class/:id", dependency.ClassHandler.GetClassById)
		}
	}

	r.Run()
}
