// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package router

import (
	"clove-api/domain/image/presentation"

	"github.com/gofiber/fiber/v2"
)

// Image Router
func NewImageRouter(image *fiber.Router, imageController *presentation.ImageController) {
	(*image).Post("/", imageController.Create)
	(*image).Get("/instance/:fileName", imageController.GetInstanceByFileName)
	(*image).Get("/", imageController.ReadAll)
	(*image).Get("/:id", imageController.Read)
	(*image).Get("/name/:name", imageController.ReadByName)
	(*image).Get("/fileName/:fileName", imageController.ReadByFileName)
	(*image).Delete("/", imageController.DeleteAll)
	(*image).Delete("/:id", imageController.Delete)
	(*image).Delete("/name/:name", imageController.DeleteByName)
	(*image).Delete("/fileName/:fileName", imageController.DeleteByFileName)
}
