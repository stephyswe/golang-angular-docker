package main

import "github.com/gofiber/fiber/v2"

func main() {
	/* dsn := "root:password@tcp(localhost:3306)/ambassador" // Replace password and adjust host/port as needed
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database")
	} */

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Listen(":8000")
}
