// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package router

import (
	"clove-api/application/router/handler"

	"github.com/gofiber/fiber/v2"
)

// Default Router
func NewDefaultRouter(app *fiber.App) {
	app.Get("/", handler.RouterOperationCheck)
	app.Get("/hello", handler.Hello)
	app.Get("/error", handler.Error)
}
