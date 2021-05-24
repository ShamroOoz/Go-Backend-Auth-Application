package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	fmt.Println("rigster controlller")
	return c.JSON({})
}

func Login(c *fiber.Ctx) error {
	fmt.Println("Login controlller")
	return nil
}

func User(c *fiber.Ctx) error {
	fmt.Println("User controlller")
	return nil
}

func Logout(c *fiber.Ctx) error {
	fmt.Println("Logout controlller")
	return nil
}