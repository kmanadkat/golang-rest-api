package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kmanadkat/go-react-tutorial/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *fiber.Ctx) error {
	// Prepare Object Id
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	// Prepare Data
	filter := bson.M{"_id": objectId}
	_, err = db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Error updating todo"})
	}

	// Respond Back
	response := fiber.Map{"msg": "Todo deleted successfully"}
	return c.Status(200).JSON(response)
}
