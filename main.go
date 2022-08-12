package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liondy/jabar-food-directory/configs"
	"github.com/liondy/jabar-food-directory/routes"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	routes.FoodsRoute(app)

	app.Listen(":8000")
}
