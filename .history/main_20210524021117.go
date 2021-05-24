package main

import (
	"./routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	
 	_, err := gorm.Open(mysql.Open("root:/go-auth"), &gorm.Config{})

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