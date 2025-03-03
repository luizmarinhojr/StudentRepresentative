package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/schema"
	"github.com/luizmarinhojr/StudentRepresentative/usecase"
)

func GetStudents(c *gin.Context) {
	students, err := usecase.GetStudents()
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student schema.StudentRequest
	err := c.BindJSON(&student)
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	uri := c.Request.RequestURI
	id, erro := usecase.CreateStudent(&student)
	if erro != nil {
		fmt.Println(erro)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.Status(http.StatusCreated)
	c.Header("location", uri+"/"+id)
}

func GetStudentById(c *gin.Context) {
	student, err := usecase.GetStudentById(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
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
