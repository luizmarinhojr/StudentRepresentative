package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/view/request"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(us usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: us,
	}
}

func (u *UserHandler) SignIn(c *gin.Context) {
	var user request.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, erro := u.usecase.SignIn(&user)
	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": erro.Error()})
		return
	}

	c.SetCookie("jwt", *token, int(time.Hour*24*7), "/", "", false, true)
	c.JSON(http.StatusAccepted, gin.H{"message": "Login succesful!"})
}

func (u *UserHandler) SignUp(c *gin.Context) {
	var login request.Login
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = u.usecase.SignUp(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusCreated)
}
