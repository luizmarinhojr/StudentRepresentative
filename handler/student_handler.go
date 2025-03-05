package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/handler/request"
	"github.com/luizmarinhojr/StudentRepresentative/usecase"
)

type StudentHandler struct {
	usecase usecase.StudentUseCase
}

func NewStudentController(us usecase.StudentUseCase) *StudentHandler {
	return &StudentHandler{
		usecase: us,
	}
}

func (sc *StudentHandler) GetStudents(c *gin.Context) {
	students, err := sc.usecase.GetStudents()
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, students)
}

func (sc *StudentHandler) CreateStudent(c *gin.Context) {
	var student request.StudentRequest
	err := c.BindJSON(&student)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = student.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uri := c.Request.RequestURI
	id, erro := sc.usecase.CreateStudent(&student)
	if erro != nil {
		log.Println(erro)
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	c.Status(http.StatusCreated)
	c.Header("location", uri+"/"+id)
}

func (sc *StudentHandler) GetStudentById(c *gin.Context) {
	student, err := sc.usecase.GetStudentById(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (sc *StudentHandler) GetOneStudentByRegistration(c *gin.Context) {
	c.JSON(200, gin.H{
		"registration": c.Param("registration"),
	})
}

func (sc *StudentHandler) GetAllStudentsByName(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name": c.Param("name"),
	})
}
