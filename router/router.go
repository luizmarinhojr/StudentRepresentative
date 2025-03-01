package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	InitializeRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
