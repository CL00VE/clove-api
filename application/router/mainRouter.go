// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package router

import (
	"clove-api/application/initializer/connector"
	"os"

	communityEnt "clove-api/domain/community/domain/ent"
	communityRouter "clove-api/domain/community/presentation/router"

	imageEnt "clove-api/domain/image/domain/ent"
	imageRouter "clove-api/domain/image/presentation/router"

	"github.com/gofiber/fiber/v2"
)

// Main App Router
func NewAppRouter(app *fiber.App, communityDB *communityEnt.Client, imageDB *imageEnt.Client) {
	v1 := app.Group(os.Getenv("SERVER_VERSION"))
	api := v1.Group("api")

	// /v1/api/post
	post := api.Group("post")
	postController := connector.PostContainerConnection(communityDB)
	communityRouter.NewPostRouter(&post, postController)
	// PostRouter Information
	// (*post).Post("/", postController.Create)
	// (*post).Get("/", postController.ReadAll) // ?order="recent or popular or comments or recommend"
	// (*post).Get("/:id", postController.Read)
	// (*post).Put("/:id/like", postController.AddLike)
	// (*post).Put("/:id", postController.Update)
	// (*post).Delete("/", postController.DeleteAll)
	// (*post).Delete("/:id", postController.Delete)

	// /v1/api/comment
	comment := api.Group("comment")
	commentController := connector.CommentContainerConnection(communityDB)
	communityRouter.NewCommentRouter(&comment, commentController)
	// CommentRouter Information
	// (*comment).Post("/", commentController.Create)
	// (*comment).Get("/", commentController.ReadAll) // ?order="recent", ?postID="{postID}"
	// (*comment).Get("/:id", commentController.Read)
	// (*comment).Put("/:id/like", commentController.AddLike)
	// (*comment).Put("/:id", commentController.Update)
	// (*comment).Delete("/", commentController.DeleteAll)
	// (*comment).Delete("/:id", commentController.Delete)

	// /v1/api/profile
	profile := api.Group("profile")
	profileController := connector.ProfileContainerConnection(communityDB)
	communityRouter.NewProfileRouter(&profile, profileController)
	// ProfileRouter Information
	// (*profile).Post("/", profileController.Create)
	// (*profile).Get("/", profileController.ReadAll)
	// (*profile).Get("/:id", profileController.Read)
	// (*profile).Put("/:id", profileController.Update)

	// /v1/api/image
	image := api.Group("image")
	imageController := connector.ImageContainerConnection(imageDB)
	imageRouter.NewImageRouter(&image, imageController)
	// ImageRouter Information
	// (*image).Post("/", imageController.Create)
	// (*image).Get("/instance/:fileName", imageController.GetInstanceByFileName)
	// (*image).Get("/", imageController.ReadAll)
	// (*image).Get("/:id", imageController.Read)
	// (*image).Get("/name/:name", imageController.ReadByName)
	// (*image).Get("/fileName/:fileName", imageController.ReadByFileName)
	// (*image).Delete("/", imageController.DeleteAll)
	// (*image).Delete("/:id", imageController.Delete)
	// (*image).Delete("/name/:name", imageController.DeleteByName)
	// (*image).Delete("/fileName/:fileName", imageController.DeleteByFileName)
}
