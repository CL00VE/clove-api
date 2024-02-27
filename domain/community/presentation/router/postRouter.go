package router

import (
	"clove-api/domain/community/presentation"

	"github.com/gofiber/fiber/v2"
)

func NewPostRouter(post *fiber.Router, postController *presentation.PostController) {
	(*post).Post("/", postController.Create)
	(*post).Get("/", postController.ReadAll) // ?order="recent or popular or comments or recommend"
	(*post).Get("/:id", postController.Read)
	(*post).Put("/:id/like", postController.AddLike)
	(*post).Put("/:id", postController.Update)
	(*post).Delete("/", postController.DeleteAll)
	(*post).Delete("/:id", postController.Delete)
}
