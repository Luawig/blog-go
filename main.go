package main

import (
	"blog-go/config"
	"blog-go/internal/db"
	"blog-go/routes"
)

func main() {
	config.InitConfig()
	db.InitDB()
	routes.InitRouter()
}
