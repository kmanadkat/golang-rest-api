package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kmanadkat/go-react-tutorial/db"
	"github.com/kmanadkat/go-react-tutorial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *fiber.Ctx) error {
	// Prepare Object Id
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo Id"})
	}

	// Prepare Todo
	todo := &models.Todo{}
	todo.ID = objectId
	err = c.BodyParser(todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid request"})
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
	}

	// Update Todo
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"body": todo.Body, "completed": todo.Completed}}
	_, err = db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Error updating todo"})
	}

	// Respond Back
	response := fiber.Map{"msg": "Todo updated successfully", "data": todo}
	return c.Status(200).JSON(response)
}
