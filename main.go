package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Todo GoLang API")
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Disconnects when server terminates
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	collection = client.Database("golang-todo").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)          // Get Todos
	app.Post("/api/todos", createTodo)       // Create Todo
	app.Patch("/api/todos/:id", updateTodo)  // Update Todo
	app.Delete("/api/todos/:id", deleteTodo) // Delete Todo

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Fatal(app.Listen(":" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	// Close Cursor After getTodos Executes
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid request"})
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
	}

	cursor, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Error creating todo"})
	}
	todo.ID = cursor.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	return nil
}

func deleteTodo(c *fiber.Ctx) error {
	return nil
}
