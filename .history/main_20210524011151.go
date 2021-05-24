package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	 dsn := "root:rootroot@/yt_go_aut"
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