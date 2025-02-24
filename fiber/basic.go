package main

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// In-memory database and ID counter
var (
	db     = make(map[int]string)
	mu     sync.Mutex // Mutex for concurrent safety
	nextID = 1        // Auto-increment ID
)

// Handler to get all items
func getItems(c *fiber.Ctx) error {
	// return *c.JSON(db) - can be as well
	return c.JSON(db)
}

// Handler to get a specific item by ID
func getItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	mu.Lock()
	item, exists := db[id]
	mu.Unlock()
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}
	return c.JSON(fiber.Map{"id": id, "value": item})
}

// Handler to add a new item with auto-increment ID
func addItem(c *fiber.Ctx) error {
	type Request struct {
		Value string `json:"value"` // The part inside backticks (json:"value") is a struct tag. This tells Go how to handle JSON serialization and deserialization.
	}
	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	mu.Lock()
	id := nextID
	nextID++
	db[id] = req.Value
	mu.Unlock()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id, "value": req.Value})
}

// Handler to delete an item by ID
func deleteItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	mu.Lock()
	_, exists := db[id]
	if exists {
		delete(db, id)
	}
	mu.Unlock()

	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.SendString("Item deleted")
}

func main() {
	app := fiber.New()

	// Routes
	app.Get("/items", getItems)
	app.Get("/items/:id", getItem)
	app.Post("/items", addItem)
	app.Delete("/items/:id", deleteItem)

	// Start server
	app.Listen(":3000")
}
