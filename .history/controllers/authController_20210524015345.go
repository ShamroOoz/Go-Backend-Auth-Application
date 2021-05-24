package controllers

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func Register(c *fiber.Ctx) error {
	fmt.Println("rigster controlller")
	return nil
}

func Login(c *fiber.Ctx) error {
	fmt.Println("rigster controlller")
	return nil
}

func User(c *fiber.Ctx) error {}

func Logout(c *fiber.Ctx) error {}