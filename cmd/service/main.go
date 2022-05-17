package main

import (
	"gogo/service/api/http"
)

func main() {
	app := http.MakeHandlers()
	app.Run()
}
