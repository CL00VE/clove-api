// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package logger

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// The NewLogger function is a function that performs the logger setting.
func NewLogger(c *fiber.Ctx) error {
	if c.Path() == os.Getenv("MONITOR_PATH") {
		return c.Next()
	}

	return logger.New(logger.Config{
		Format: fmt.Sprintf("%s\n", os.Getenv("LOGGER_FORMAT")),
	})(c)
}
