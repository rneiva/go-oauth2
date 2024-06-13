package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-oauth2/controllers"
)

func main() {
	app := fiber.New()

	app.Post("/google_login", controllers.GoogleLogin)
	//app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8000")
}
