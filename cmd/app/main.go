package main

import (
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/internal/database"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/router"
)

func main() {
	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}

	router.InitializeApi(db)
}
