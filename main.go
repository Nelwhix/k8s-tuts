package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT = "3000"

func main() {
	http.HandleFunc("/", HostHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %v\n", err)
	}

	output := fmt.Sprintf("Hello from host: %s. \n Current time is: %s. \n", hostname, time.Now().String())
	_, err = fmt.Fprintf(w, output)
	if err != nil {
		log.Printf("Error printing to response: %v\n", err)
	}
}
