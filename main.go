package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-oauth2/config"
	"github.com/rneiva/go-oauth2/controllers"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Post("/google_login", controllers.GoogleLogin)
	//app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8000")
}
