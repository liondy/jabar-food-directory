package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/liondy/jabar-food-directory/api/configs"
	"github.com/liondy/jabar-food-directory/api/models"
	"github.com/liondy/jabar-food-directory/api/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = configs.GetCollection(configs.DB, "foods")

func GetAllFoods(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// foodCollection.Find(ctx, bson.D{})
	var foods []models.Food
	defer cancel()

	results, err := foodCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BaseResponse{Status: http.StatusInternalServerError, Message: "error 1", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleFood models.Food
		if err = results.Decode(&singleFood); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.BaseResponse{Status: http.StatusInternalServerError, Message: "error food", Data: &fiber.Map{"data": err.Error()}})
		}

		foods = append(foods, singleFood)
	}

	return c.Status(http.StatusOK).JSON(
		responses.BaseResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": foods}},
	)
}
