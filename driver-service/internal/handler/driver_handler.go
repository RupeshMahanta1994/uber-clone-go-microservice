package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rupeshmahanta/driver-service/internal/model"
	"github.com/rupeshmahanta/driver-service/internal/service"
)

type DriverHandler struct {
	service service.IDriverService
}

func NewDriverHandler(service service.IDriverService) *DriverHandler {
	return &DriverHandler{service: service}
}

func (h *DriverHandler) RegisterDriverRoutes(app *fiber.App) {
	app.Post("/drivers/onboard", h.Onboard)
	app.Put("/drivers/availability", h.ToggleAvailability)
	// app.Put("/drivers/status", h.UpdateStatus)
	app.Get("/drivers/me", h.GetDriver)
}
func (h *DriverHandler) Onboard(c *fiber.Ctx) error {
	userId := c.Get("X-User-ID")
	var req OnboardRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error in parsing the Onboarding details")
		log.Print(string(c.Body()))
		return errors.New("Error in parsing Onboarding Info")
	}
	driver := &model.Driver{
		UserID:        userId,
		VehicleNumber: req.VehicleNumber,
		VehicleType:   req.VehicleType,
	}
	if _, err := h.service.Onboard(c.Context(), driver); err != nil {
		log.Println("Error in Onboarding Driver")
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}
	log.Println("Onboarding Driver is completed")
	return c.JSON(fiber.Map{"message": "Driver onboarded"})
}
func (h *DriverHandler) ToggleAvailability(c *fiber.Ctx) error {
	userId := c.Get("X-User-ID")
	var req ToggleAvailabilityRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error in parsing Driver availability")
		return c.Status(500).JSON(fiber.Map{"Error": "Parsing Availability payload"})
	}
	if _, err := h.service.ToggleAvailability(c.Context(), userId, req.Availability); err != nil {
		log.Println("Error in changing availability status of the driver")
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}
	log.Println("Availability status changed successfully")
	return c.JSON(fiber.Map{"message": "Availability status changed successfully"})
}

func (h *DriverHandler) GetDriver(c *fiber.Ctx) error {
	userId := c.Get("X-User-ID")

	driver, err := h.service.GetDriver(c.Context(), userId)
	if err != nil {
		log.Println("Error in getting Driver details Handler")
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}
	log.Println("Successfully retreived Driver data")
	return c.Status(200).JSON(fiber.Map{"message": "Driver details", "Driver data": driver})
}
