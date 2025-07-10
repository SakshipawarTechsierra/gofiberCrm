package lead

import (
	"github.com/SakshipawarTechsierra/gofiberCrm/database"
	"github.com/gofiber/fiber/v2"
)

// Just use database.Lead instead of importing it
func GetLeads(c *fiber.Ctx) error {
	var leads []database.Lead
	database.DBConn.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead database.Lead
	if err := database.DBConn.First(&lead, id).Error; err != nil {
		return c.Status(404).SendString("Lead not found")
	}
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	lead := new(database.Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DBConn.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead database.Lead
	if err := database.DBConn.First(&lead, id).Error; err != nil {
		return c.Status(404).SendString("Lead not found")
	}
	database.DBConn.Delete(&lead)
	return c.SendString("Lead deleted successfully")
}
