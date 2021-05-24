package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
   return c.SendString("rigster controlller 👋!")
}

func Login(c *fiber.Ctx) error {
return c.SendString("Login controlller")
	
}

func User(c *fiber.Ctx) error {
	return c.SendString("User controlller")
	return nil
}

func Logout(c *fiber.Ctx) error {
	fmt.Println("Logout controlller")
	return nil
}