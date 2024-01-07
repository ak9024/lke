package server

import "github.com/gofiber/fiber/v2"

func Router() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(struct {
			Message string `json:"message"`
		}{
			Message: "ok!",
		})
	})

	return app
}

func StartApp() error {
	server := Router()

	if err := server.Listen(":4000"); err != nil {
		return err
	}

	return nil
}
