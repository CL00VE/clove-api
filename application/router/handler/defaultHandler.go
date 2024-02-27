package handler

import (
	"clove-api/global/exception/global"

	"github.com/gofiber/fiber/v2"
)

// The RouterOperationCheck function performs the task of verifying that the router is operating normally.
func RouterOperationCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Success")
}

// Result Test
func Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Hello!")
}

// Error Test
func Error(c *fiber.Ctx) error {
	global.NewTestException("original err message").Throw()
	return nil
}
