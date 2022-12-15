package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	
	database.GetConexao()
	routes.HandleResquest()
}
