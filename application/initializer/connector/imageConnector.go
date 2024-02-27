// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package connector

import (
	"clove-api/domain/image/domain/ent"
	"clove-api/domain/image/domain/repository"
	"clove-api/domain/image/presentation"
	"clove-api/domain/image/service"
)

// The ImageContainerConnection function injects the dependence of services related to Image objects.
func ImageContainerConnection(db *ent.Client) *presentation.ImageController {
	imageRepository := repository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepository)
	imageController := presentation.NewImageController(imageService)
	return imageController
}
