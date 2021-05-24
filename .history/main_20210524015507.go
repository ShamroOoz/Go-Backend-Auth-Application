package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	 
		app := fiber.New()

		

		app.Listen(":4000")
	
}