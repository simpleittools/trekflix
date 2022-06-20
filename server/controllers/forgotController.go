package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/trekflix/database"
	"github.com/simpleittools/trekflix/models"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	token := RandomStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := os.Getenv("PASSWORDRESETFROM")
	to := []string{
		data["email"],
	}
	url := os.Getenv("FRONTEND") + "/reset/" + token
	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")

	err = smtp.SendMail(os.Getenv("SMTPSERVER")+os.Getenv("SMTPPORT"), nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

// RandomStringRunes generate a random string. Found at https://golangdocs.com/generate-random-string-in-golang
func RandomStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	var passwordReset = models.PasswordReset{}

	if err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", hashedPassword)

	// Remove the token from the database, so it cannot be reused
	database.DB.Model(&models.PasswordReset{}).Where("token = ?", data["token"]).Delete(&data)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
