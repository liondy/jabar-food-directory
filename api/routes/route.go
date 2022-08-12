package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liondy/jabar-food-directory/api/controllers"
)

func FoodsRoute(app *fiber.App) {
	app.Get("/foods", controllers.GetAllFoods)
}
