package main

import (
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/repository"
	"github.com/luizmarinhojr/StudentRepresentative/router"
)

func main() {
	db, err := repository.OpenConnection()
	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}

	router.InitializeApi(db)
}
