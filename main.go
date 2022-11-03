package main

import (
	"fuzzygolang/internal/app"
)

func main() {
	app := app.StartApp()
	app.Run()
}
