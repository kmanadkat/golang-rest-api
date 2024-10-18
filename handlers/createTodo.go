package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kmanadkat/go-react-tutorial/db"
	"github.com/kmanadkat/go-react-tutorial/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid request"})
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
	}

	cursor, err := db.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Error creating todo"})
	}
	todo.ID = cursor.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}
