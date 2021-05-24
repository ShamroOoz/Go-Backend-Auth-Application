package main

import (
	"go-auth-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	
 	//database.Connect()
	_, err :=gorm.Open(mysql.Open("root:@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}