package router

import (
	"clove-api/domain/community/presentation"

	"github.com/gofiber/fiber/v2"
)

func NewCommentRouter(comment *fiber.Router, commentController *presentation.CommentController) {
	(*comment).Post("/", commentController.Create)
	(*comment).Get("/", commentController.ReadAll) // ?order="recent", ?postID="{postID}"
	(*comment).Get("/:id", commentController.Read)
	(*comment).Put("/:id/like", commentController.AddLike)
	(*comment).Put("/:id", commentController.Update)
	(*comment).Delete("/", commentController.DeleteAll)
	(*comment).Delete("/:id", commentController.Delete)
}
