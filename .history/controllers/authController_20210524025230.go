package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
    var data map[string]type
   return c.SendString("rigster controlller 👋!")
}

func Login(c *fiber.Ctx) error {
return c.SendString("Login controlller")
	
}

func User(c *fiber.Ctx) error {
	return c.SendString("User controlller")
	
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout controlller")
	
}