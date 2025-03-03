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

	// // Teste aqui
	// st := schema.Student{
	// 	Name: "João",
	// }

	// r := reflect.ValueOf(&st).Elem()

	// numFields := r.NumField()

	// v := r.Field(0).Addr().String()

	// fmt.Println(r)
	// fmt.Println(numFields)

	// fmt.Println(v)
}
