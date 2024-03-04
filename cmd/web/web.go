package main

import (
	"os"
	"os/signal"
	"strconv"

	"github.com/krishan1988/home24-web-analyser/internal/api"
)

func main() {
	osSig := make(chan os.Signal, 1)

	signal.Notify(osSig, os.Interrupt)

	portStr := os.Getenv("SERVICE_PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port =	8080
	}

	api.Configure(port)

	<-osSig
}