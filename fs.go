package main

import (
	"log"
	"net/http"
	"time"
)

const (
	ADDR        = "localhost:8083"
	FILED       = "static"
	TIMEOUT_SEC = 30
)

func main() {
	// Create new server
	server := &http.Server{
		// Listen only to localhost
		Addr: ADDR,
		// Direct any request to static folder
		Handler: http.FileServer(http.Dir(FILED)),
		// Set timeout, because normally none and overflow
		ReadTimeout: TIMEOUT_SEC * time.Second,
	}
	log.Printf("fs: addr=%s; dir=%s; timeout=%d sec\n", ADDR, FILED, TIMEOUT_SEC)
	err := server.ListenAndServe()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
