package controllers

import (
	"go-auth-app/database"
	"go-auth-app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string
	
   if error :=c.BodyParser(&data); error !=nil {
		return error
	}

	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
	Name:   data["name"] ,
	Email: data["email"] ,
	Password: password,
	}

   database.DB.Create(&user)
   
   return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	
}

func User(c *fiber.Ctx) error {
	return c.SendString("User controlller")
	
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout controlller")
	
}