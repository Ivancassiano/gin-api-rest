package main

import (
	"github.com/Ivancassiano/gin-api-rest/database"
	"github.com/Ivancassiano/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
