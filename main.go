package main

import (
	"order_kafe/config"
	"order_kafe/database"
)

func main() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
}
