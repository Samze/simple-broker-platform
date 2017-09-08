package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/samze/simple-broker-platform/handlers"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to start on")
	flag.Parse()

	startServer(port)
}

func startServer(port int) {
	address := fmt.Sprintf("127.0.0.1:%d", port)

	handler := handlers.New()

	srv := &http.Server{
		Handler:      handler,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
