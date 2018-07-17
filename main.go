package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	startTime time.Time
)

func init() {
	startTime = time.Now()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, err := os.Hostname()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Hostname: %v\nUptime: %v\n", name, time.Since(startTime))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
