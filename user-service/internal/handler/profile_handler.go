package handler

import (
	"context"
	"github/rupeshmahanta/user-service/internal/models"
	"github/rupeshmahanta/user-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	service service.IProfileService
}

func NewProfileHandler(service service.IProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) RegisterProfileRoutes(app *fiber.App) {
	app.Post("/profile/create", h.CreateProfile)
	app.Get("/profile/:userId", h.GetUserById)
	app.Put("/profile/update", h.UpdateUser)

}

func (h *ProfileHandler) CreateProfile(c *fiber.Ctx) error {
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	// pass a context and a pointer to the profile
	userData, err := h.service.CreateProfile(context.Background(), &profile)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "profile created", "User Data": userData})
}
func (h *ProfileHandler) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	profile, err := h.service.GetProfileById(c.Context(), userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "profile retrieved", "User Data": profile})

}
func (h *ProfileHandler) UpdateUser(c *fiber.Ctx) error {
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Invalid Input"})
	}
	err := h.service.UpdateProfile(c.Context(), &profile)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Profile updated successfully"})

}
