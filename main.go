package main

import (
	"github.com/sid-shaha/restAPIExample/server"
)

func main() {
	app := server.NewApp()
	app.Run()
}
