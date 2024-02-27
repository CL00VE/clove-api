// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package response

// General Response
type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"messgae"`
	Data    any    `json:"data,omitempty"`
}
