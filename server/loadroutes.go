package server

import "github.com/Dank-del/link-shortner-api/server/routes"

func loadRoutes() {
	ServerObject.GET("link", routes.GetRedirect)
	ServerObject.GET("new", routes.Shorten)
}
