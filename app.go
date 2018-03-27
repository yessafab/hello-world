package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// initial config of environment variable(s)
	port := os.Getenv("TELSTRA-HELLOWORLD_PORT")
	if port == "" {
		port = "3333"
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)

	// handlers
	http.HandleFunc("/", rootHandle)

	// listen to address and fatal if we ever crash as this should be up
	// permanently
	log.Printf("Hello World Listening locally on - %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
