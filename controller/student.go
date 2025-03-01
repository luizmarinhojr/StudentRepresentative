package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/repository"
)

func GetAllStudents(c *gin.Context) {
	students, err := repository.GetAllStudents()
	if err != nil {
		panic("get all students doesn't work")
	}
	c.JSON(200, students)
}

func CreateStudent(c *gin.Context) {

}

func GetOneStudentByRegistration(c *gin.Context) {
	c.JSON(200, gin.H{
		"registration": c.Param("registration"),
	})
}

func GetAllStudentsByName(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name": c.Param("name"),
	})
}
