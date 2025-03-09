package main

import (
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/dependencies"
	"github.com/luizmarinhojr/StudentRepresentative/internal/database"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/router"
)

func main() {
	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}

	dependencies := dependencies.Inject(db)

	router.InitializeApi(*dependencies)
}
