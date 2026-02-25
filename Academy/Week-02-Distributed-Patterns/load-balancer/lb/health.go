package main

import (
	"log"
	"net/http"
	"time"
)

func isHealthy(server string) bool {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(server + "/health")
	if err != nil {
		log.Printf("❌ [%s] is not healthy\n", server)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ [%s] is not healthy\n", server)
		return false
	}
	return true
}
