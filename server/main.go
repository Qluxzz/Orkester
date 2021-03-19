package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/stream", func(c *fiber.Ctx) error {
		c.Response().Header.Add("content-type", "audio/ogg")

		r, err := os.Open("./content/test.ogg")
		if err != nil {
			panic(err)
		}

		return c.SendStream(r)
	})

	app.Listen(":3000")
}
