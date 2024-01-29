package main

import (
	"github.com/gin-gonic/gin"
	"project.com/API/db"
	"project.com/API/routes"
)
func main(){
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}