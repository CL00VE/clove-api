// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package monitor

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// The NewLogger function is a function that performs the monitor setting.
func NewMonitor(c *fiber.Ctx) error {
	return monitor.New(monitor.Config{
		Title: fmt.Sprintf("%s Server Monitor", os.Getenv("SERVER_NAME")),
	})(c)
}
