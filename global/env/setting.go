// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package env

import (
	"clove-api/global/enum"

	"github.com/joho/godotenv"
)

// The EnvFileLoad function is a function that automatically loads all Env files when the server is running to help smooth server operation.
func EnvFileLoad() error {
	envFilePaths := []string{
		"../resource/app.env",
		"../resource/database.env",
		"../resource/server.env",
		"../resource/status.env",
	}

	if err := godotenv.Load(envFilePaths...); err != nil {
		return err
	}

	envLoadNext()

	return nil
}

// env load next
func envLoadNext() {
	enum.ErrorTypeSetting()
}
