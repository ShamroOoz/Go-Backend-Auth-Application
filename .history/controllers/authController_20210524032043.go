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
   
   return c.JSON(userINSERT INTO users (id, name, email, password)
   VALUES (
	   'id:bigint',
	   'name:longtext',
	   'email:varchar',
	   'password:longblob'
	 );)
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