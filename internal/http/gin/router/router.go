package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/dependencies"
)

func InitializeApi(dependency dependencies.Dependencies) {
	r := gin.Default()
	InitializeRoutes(r, dependency)
}
