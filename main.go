package main

import (
	"github.com/Dank-del/link-shortner-api/database"
	"github.com/Dank-del/link-shortner-api/server"
	"github.com/Dank-del/link-shortner-api/utils"
)

func main() {
	database.StartDatabase()
	server.Start(utils.GetConfig())
}
