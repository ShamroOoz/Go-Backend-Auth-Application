package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	 _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
		if err != nil {
			panic(err)
		}

		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World 👋!")
		})

		app.Listen(":4000")
	
}