// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9

// The package main is an API server built on the Fiber framework.
// In consideration of maintenance, the project file structure was designed in a domain driving design (DDD) manner.
// It was developed with easy API server management in mind.
package main

import (
	"clove-api/application/filter/cors"
	"clove-api/application/filter/logger"
	"clove-api/application/filter/monitor"
	"clove-api/application/initializer"
	"clove-api/application/middleware"
	"clove-api/application/router"
	"clove-api/application/status"
	"clove-api/global/env"
	"clove-api/global/exception/handler"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// A Port variable is a variable that stores the server port as a string type.
var port string

func init() {
	if err := env.EnvFileLoad(); err != nil {
		panic(err)
	}

	if err := status.Setting(); err != nil {
		panic(err)
	}

	flogPort := flag.String("p", os.Getenv("SERVER_PORT"), "Enter the port")
	flag.Parse()
	port = fmt.Sprintf(":%s", *flogPort)
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:      os.Getenv("SERVER_NAME"),
		Prefork:      true,
		ErrorHandler: handler.ExceptionHandler,
	})

	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.NewCors)
	app.Use(logger.NewLogger)
	app.Use(os.Getenv("MONITOR_PATH"), monitor.NewMonitor)

	db := initializer.NewDatabase()

	communityDB, err := db.NewCommunityDatabase()
	if err != nil {
		panic(err)
	}

	imageDB, err := db.NewImageDatabase()
	if err != nil {
		panic(err)
	}

	router.NewDefaultRouter(app)
	router.NewAppRouter(app, communityDB, imageDB)

	app.Use(middleware.NotFoundHandler)

	log.Fatal(app.Listen(port))
}
