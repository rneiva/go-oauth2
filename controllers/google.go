package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-oauth2/config"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)

	return c.JSON(url)

}
