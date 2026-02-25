package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHashRing(t *testing.T) {
	ring := New(3)
	ring.Add("node1")
	ring.Add("node2")
	ring.Add("node3")

	node := ring.Get("user123")
	for i := 0; i < 10; i++ {
		if node != ring.Get("user123") {
			t.Error("data mismatch")
		}
	}
}

func TestMockHTTPServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	if !isHealthy(server.URL) {
		t.Error("server should be healthy")
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	if isHealthy(server.URL) {
		t.Error("server should be busy")
	}
}
