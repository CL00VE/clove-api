// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package middleware

import (
	"clove-api/global/exception/system"

	"github.com/gofiber/fiber/v2"
)

// The NotFoundHandler function is a handler that handles NotFound errors.
func NotFoundHandler(c *fiber.Ctx) error {
	system.NewNotFoundException(c.OriginalURL()).Throw()
	return nil
}
