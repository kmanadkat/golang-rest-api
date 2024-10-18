package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kmanadkat/go-react-tutorial/db"
	"github.com/kmanadkat/go-react-tutorial/handlers"
)

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
	db.Collection = client.Database("golang-todo").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", handlers.GetTodos)          // Get Todos
	app.Post("/api/todos", handlers.CreateTodo)       // Create Todo
	app.Patch("/api/todos/:id", handlers.UpdateTodo)  // Update Todo
	app.Delete("/api/todos/:id", handlers.DeleteTodo) // Delete Todo

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Fatal(app.Listen(":" + port))
}
