package main

import (
	"log"

	"todoAPI/db"
	"todoAPI/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err:= godotenv.Load()
	if err!=nil{
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.Connect()





	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
