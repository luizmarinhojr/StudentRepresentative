package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase"
)

type ClassHandler struct {
	usecase usecase.ClassUseCase
}

func NewClassHandler(us usecase.ClassUseCase) *ClassHandler {
	return &ClassHandler{
		usecase: us,
	}
}

func (ch *ClassHandler) GetClasses(ctx *gin.Context) {
	classes, err := ch.usecase.GetClasses()
	if err != nil {
		fmt.Println(err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, classes)
}

func (ch *ClassHandler) GetClassById(ctx *gin.Context) {
	class, err := ch.usecase.GetClassById(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, class)
}
