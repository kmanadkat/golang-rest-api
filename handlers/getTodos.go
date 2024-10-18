package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kmanadkat/go-react-tutorial/db"
	"github.com/kmanadkat/go-react-tutorial/models"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	// Close Cursor After getTodos Executes
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}
