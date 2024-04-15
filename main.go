package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/getMe", func(c *fiber.Ctx) error {
		agent := fiber.Get("https://httpbin.org/get")
		statusCode, body, errs := agent.Bytes()
		if len(errs) > 0 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errs": errs,
			})
		}

		var resData fiber.Map
		err := json.Unmarshal(body, &resData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}

		return c.Status(statusCode).JSON(resData)
	})

	// Start server
	app.Listen(":3000")
}
