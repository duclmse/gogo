package http

import (
	"fmt"
	"time"

	"github.com/aerogo/aero"
)

func storeSession(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		// Handle the request first.
		err := next(ctx)
		// If the session was modified, store it.
		if ctx.HasSession() && ctx.Session().Modified() {
			err := ctx.App().Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
			if err != nil {
				fmt.Printf("%s", err)
				return err
			}
		}
		return err
	}
}

func elapse(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		// Measure response time
		start := time.Now()
		err := next(ctx)
		// Write it to the log
		fmt.Printf("%s\n", time.Since(start))
		// Make sure to pass the error back!
		return err
	}
}
