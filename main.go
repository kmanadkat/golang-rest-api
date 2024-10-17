package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Todo GoLang API")
  // Load Environment Variables
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	app := fiber.New()

	// Data
	todos := []Todo{}

	// Handlers
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})

	// Create a Todo
	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		// Validation
		err := c.BodyParser(todo)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "Invalid request"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
		}

		// Form Todo
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		// Respond Back
		return c.Status(201).JSON(todo)
	})

	// Update a Todo
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		todo := &Todo{}
		// Validation
		err := c.BodyParser(todo)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": "Invalid request"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
		}

		// Find Todo
		for i, t := range todos {
			if fmt.Sprint(t.ID) == id {
				todos[i].Body = todo.Body
				todos[i].Completed = todo.Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	})

	// Get All Todos
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

  // Delete a Todo
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Find Todo
		for i, t := range todos {
			if fmt.Sprint(t.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"sucess": true})
			}
		}
		
		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	})

  PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":"+PORT))
}