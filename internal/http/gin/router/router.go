package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeApi(db *sql.DB) {
	r := gin.Default()
	InitializeRoutes(r, db)
}
