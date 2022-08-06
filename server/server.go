package server

import (
	"fmt"
	"github.com/Dank-del/link-shortner-api/types"
	"github.com/labstack/echo/v4"
)

func Start(conf *types.Config) {
	ServerObject = echo.New()
	loadRoutes()
	ServerObject.Logger.Fatal(
		ServerObject.Start(fmt.Sprintf("%s:%d", conf.Host, conf.Port)),
	)
}
