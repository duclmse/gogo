package http

import (
	"fmt"

	"github.com/aerogo/aero"
)

func MakeHandlers() *aero.Application {
	app := aero.New()
	//routing
	app.Get("/", home)

	app.Post("/hello", hello)
	app.Get("/hello/:person", helloPerson)

	app.Get("/images/*file", imageFile)
	app.Get("/streamhello", streamHello)

	//middleware
	app.Use(storeSession)
	app.Use(elapse)
	//app.Use(one, two, three)

	app.OnStart(func() {
		fmt.Printf("Started server\n")
	})
	app.OnEnd(func() {
		fmt.Printf("Stop server\n")
	})

	return app
}
