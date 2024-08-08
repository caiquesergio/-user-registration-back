package main

import (
	"github.com/caiquesergio/api-go-gin/database"
	"github.com/caiquesergio/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
