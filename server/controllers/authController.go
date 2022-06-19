package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/trekflix/database"
	"github.com/simpleittools/trekflix/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func Register(c *fiber.Ctx) error {
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

	// hash the password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Username:  data["username"],
		Email:     data["email"],
		Password:  hashPassword,
	}

	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User

	// Search the Database for the user by email
	database.DB.Where("username = ?", data["username"]).First(&user)
	// TODO: Do Not Deploy in this state, Need to hide which fails between the username and password
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	//Generate JWT Token

	claims := jwt.StandardClaims{
		// convert the unsigned integer of user.ID to a string
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// todo: Set a Secret and replace "secret" with a new secret on each use
	token, err := jwtToken.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Create Cookie
	cookie := fiber.Cookie{
		Name:     "trekflix",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"trekflix": token,
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("trekflix")

	//	parse the token
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*Claims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "trekflix",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
