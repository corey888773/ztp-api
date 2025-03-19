package main

import (
	"context"
	"log"

	"github.com/corey888773/ztp-api/src/app"
	"github.com/corey888773/ztp-api/src/util"
)

func main() {
	// Create context
	appCtx, appCancel := context.WithCancel(context.Background())
	go func() {
		<-appCtx.Done()
		log.Println("Shutting down the application in 2 seconds...")
	}()

	// Start API server
	server, err := app.CreateApp(appCtx)
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}

	err = server.Start(":8000")
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}
}
