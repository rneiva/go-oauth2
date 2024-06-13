package controllers

import (
	"context"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/rneiva/go-oauth2/config"
)

func GithubLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GithubLoginConfig.AuthCodeURL("randomstate")
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)

	return c.JSON(url)
}

func GithubCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't match")
	}

	code := c.Query("code")
	githuconn := config.GithubConfig()

	token, err := githuconn.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token exchange failed")
	}
	fmt.Println(token)

	resp, err := githuconn.Client(context.Background(), token).Get("https://api.github.com/repo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data fetch failed")
	}
	fmt.Println(resp)

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))
}
