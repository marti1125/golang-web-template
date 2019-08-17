package main

import (
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func server() {
	go func() {
		http.ListenAndServe(":8080", Route())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	log.Info("Running...")
	<-c
}

func main() {
	server()
}
