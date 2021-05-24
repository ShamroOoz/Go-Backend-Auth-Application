package controllers

import (
	"go-auth-app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string
	
   if error :=c.BodyParser(&data); error !=nil {
		return error
	}

	password,_ := bcrypt.GenerateFromPassword(data["password"], cost:14)

	user := models.User{
	Name:   data["name"] ,
	Email: data["email"] ,
	Password: password
	}

   return c.JSON(data)
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