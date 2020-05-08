package main

import (
	Config "ttemplate/conf"
	"ttemplate/server"
)

func main() {
	Config.LoadAndInitConfig()

	r := server.NewRouter()

	r.Run(":8000")
}
