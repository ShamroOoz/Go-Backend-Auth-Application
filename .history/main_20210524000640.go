package main



import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":4000")
}