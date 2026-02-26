package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s\n", r.RemoteAddr)
		log.Printf("User-ID: %v\n", r.URL.Query().Get("user_id"))
		hostName, _ := os.Hostname()
		fmt.Fprintf(w, "[%s] Hello from backend server on port %d\n", hostName, *port)
	})

	log.Printf("Starting backend server on port %d\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
