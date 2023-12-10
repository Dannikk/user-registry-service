package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"user_registry/internal/app"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			log.Println("Oh oh oh, ERROR:", e)
		}
	}()

	app, err := app.NewApp(".env")

	if err != nil {
		log.Panicf("Error: %v\n", err)
	}

	go func() {
		log.Println(app.StartHTTPServer())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Printf("%v was called. Shutdown the app", sig)
}
