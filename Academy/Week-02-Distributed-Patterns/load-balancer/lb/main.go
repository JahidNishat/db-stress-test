package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"slices"
	"time"
)

func (ch *ConsistentHash) processLBReq(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	proxyStr := ch.Get(userID)
	if proxyStr == "" {
		http.Error(w, "no available backends", http.StatusServiceUnavailable)
		return
	}
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Printf("url parse error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	ring := New(50)
	serviceName := os.Getenv("BACKENDS_SERVICE")
	if serviceName == "" {
		ring.Add("http://backend1:8080")
		ring.Add("http://backend2:8080")
		ring.Add("http://backend3:8080")
		ring.Add("http://backend4:8080")
	} else {
		go func() {
			for {
				time.Sleep(5 * time.Second)

				// Chunk 1: Discovery
				ips, err := net.LookupHost(serviceName)
				if err != nil {
					log.Println("dns look up error: ", err)
					continue
				}

				// Chunk 2: Health Check
				var alive []string
				for _, ip := range ips {
					addr := fmt.Sprintf("http://%s:8080", ip)
					if isHealthy(addr) {
						alive = append(alive, addr)
					}
				}

				// Chunk 3: Remove Dead Node
				nodes := ring.Nodes()
				for _, node := range nodes {
					if !slices.Contains(alive, node) {
						log.Println("Removing dead node: ", node)
						ring.Remove(node)
					}
				}

				// Chunk 4: Add new nodes
				for _, node := range alive {
					if !ring.Contains(node) {
						log.Println("Adding new node: ", node)
						ring.Add(node)
					}
				}
			}
		}()
	}

	http.HandleFunc("/", ring.processLBReq)
	log.Println("Starting LB on port :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
