package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/api/v1/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "halo selamat datang di api go halah"})
	})

	log.Fatal(app.Listen(":3000"))
}
