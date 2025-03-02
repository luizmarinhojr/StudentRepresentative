package main

import (
	"log"

	"github.com/luizmarinhojr/StudentRepresentative/repository"
	"github.com/luizmarinhojr/StudentRepresentative/router"
)

func main() {
	_, err := repository.OpenConnection(true)
	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}
	router.Initialize()
}
