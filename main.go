package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-oauth2/config"
	"github.com/rneiva/go-oauth2/controllers"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()
	config.GithubConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)

	app.Get("/github_login", controllers.GithubLogin)
	app.Get("/github_callback", controllers.GithubCallback)

	app.Listen(":8000")
}
