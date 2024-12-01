package main

import (
	"fmt"
	"rentalMobil/internal/routes"
	"rentalMobil/pkg/database"
	"rentalMobil/pkg/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	db := database.ConnectToMysql()
	if db != nil {
		fmt.Println("connected")
	}

	r.GET("/ping", middleware.AuthJWT(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes
	routes.UserRoutes(r, db)
	routes.CarsRoutes(r, db)

	r.Run()
}
