package main

import (
  "github.com/gofiber/fiber/v2"
 "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {

	 _, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "my_mysql_driver",
		DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", 
		}), &gorm.Config{})
		if err != nil {
	panic(err)
}else{
	
}
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":4000")
	
}