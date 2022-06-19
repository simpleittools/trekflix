package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/trekflix/controllers"
)

func Setup(app *fiber.App) {
	//app.Get("/", controllers.Index)

	auth := app.Group("api/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
