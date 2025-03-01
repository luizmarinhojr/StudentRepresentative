package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/controller"
)

func InitializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	fromBase := r.Group(basePath)
	{
		{
			fromBase.GET("/students", controller.GetAllStudents)
			fromBase.POST("student", controller.CreateStudent)
			fromBase.GET("/students/name/:name", controller.GetAllStudentsByName)
			fromBase.GET("/students/registration/:registration", controller.GetOneStudentByRegistration)
		}
	}
}
