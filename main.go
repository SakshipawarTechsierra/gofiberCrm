package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/SakshipawarTechsierra/gofiberCrm/database"
	"github.com/SakshipawarTechsierra/gofiberCrm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()

	// Initialize the database
	database.Connect()

	fmt.Println("✅ Database connected and migrated")

	// Setup routes
	setupRoutes(app)

	// Start server
	app.Listen(":3000")
}
