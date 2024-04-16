package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		
		resp, err := http.Get("https://httpbin.org/get")
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error making GET request")
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error reading response body")
		}
		//fmt.Println(resp.StatusCode)

		return c.SendString(string(body))
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

	app.Post("/postMe", func(c *fiber.Ctx) error {

		agent := fiber.Post("https://httpbin.org/post")
		agent.Body(c.Body()) // set body

		statusCode, body, errs := agent.Bytes()
		if len(errs) > 0 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errs": errs,
			})
		}

		return c.Status(statusCode).Send(body)
	})
	app.Post("/postWithoutdata", func(c *fiber.Ctx) error {

		resp, err := http.Post("https://httpbin.org/post", "application/json", nil)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error making POST request")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error reading response body")
		}

		return c.SendString(string(body))
	})

	// Start server
	app.Listen(":3000")
}
