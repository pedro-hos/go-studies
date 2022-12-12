package main

import (
	"fmt"
	"go-studies/alura-course/go-api/api-go-rest/database"
	"go-studies/alura-course/go-api/api-go-rest/routes"
)

func main() {
	fmt.Println("Iniciando Server Rest Go")
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
