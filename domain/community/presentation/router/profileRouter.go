package router

import (
	"clove-api/domain/community/presentation"

	"github.com/gofiber/fiber/v2"
)

func NewProfileRouter(profile *fiber.Router, profileController *presentation.ProfileController) {
	(*profile).Post("/", profileController.Create)
	(*profile).Get("/", profileController.ReadAll)
	(*profile).Get("/:id", profileController.Read)
	(*profile).Put("/:id", profileController.Update)
}
