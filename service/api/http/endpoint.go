package http

import (
	"fmt"
	"io"

	"github.com/aerogo/aero"
	"github.com/goccy/go-json"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

func home(ctx aero.Context) error {
	// Load number of views
	views := 0
	storedViews := ctx.Session().Get("views")
	if storedViews != nil {
		views = storedViews.(int)
	}
	views++                                         // Increment
	ctx.Session().Set("views", views)               // Store number of views
	return ctx.Text(fmt.Sprintf("%d views", views)) // Display current number of views
}

type jsonName struct {
	Name string `json:"name"`
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func hello(ctx aero.Context) error {
	req := ctx.Request()
	res := ctx.Response().Internal()
	if req.Header(contentType) != applicationJSON {
		res.WriteHeader(400)
		return JSON(ctx, response{Code: 1, Msg: "invalid type, expect application/json"})
	}
	name := jsonName{}
	bytes, err := req.Body().Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &name)
	if err != nil {
		return err
	}
	return JSON(ctx, response{Code: 0, Msg: fmt.Sprintf("Hello %s", name.Name)})
}

func helloPerson(ctx aero.Context) error {
	return ctx.String("Hello " + ctx.Get("person"))
}

func imageFile(ctx aero.Context) error {
	return ctx.String(ctx.Get("file"))
}

func streamHello(ctx aero.Context) error {
	reader, writer := io.Pipe()

	go func() {
		for i := 0; i < 100000; i++ {
			_, err := writer.Write([]byte("Hello\n"))
			if err != nil {
				fmt.Printf("Error writting response %s\n", err.Error())
				err := writer.CloseWithError(err)
				if err != nil {
					fmt.Printf("Error closing writer %s\n", err.Error())
				}
				return
			}
		}

		if err := writer.Close(); err != nil {
			fmt.Printf("Error closing writer %s\n", err.Error())
			return
		}
	}()

	return ctx.Reader(reader)
}

func JSON(ctx aero.Context, value interface{}) error {
	ctx.Response().SetHeader(contentType, applicationJSON)
	bytes, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return ctx.Bytes(bytes)
}
