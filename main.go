package main

import (
	"order_kafe/config"
	"order_kafe/database"
	"order_kafe/routes"
)

func main() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
	db := database.DB

	routes.ApiRoutes(db)
}
