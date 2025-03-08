package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/auth"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
		log.Println("no token JWT available on request")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Printf("token JWT: %v", token)

	claims, erro := auth.ValidateJWT(&token)
	if erro != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user_id", claims.UserID)
	log.Printf("user_id: %v", claims.UserID)
	c.Next()
}
