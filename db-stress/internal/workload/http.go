package workload

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type HTTPWorkload struct {
	URL    string
	Client *http.Client
}

func (w *HTTPWorkload) Name() string {
	return "HTTP Load Balancer Test"
}

func (w *HTTPWorkload) Setup(db *sql.DB) error {
	w.Client = &http.Client{
		Timeout: 2 * time.Second,
	}
	return nil
}

func (w *HTTPWorkload) Run(db *sql.DB) error {
	id := rand.Intn(1000)
	w.URL = fmt.Sprintf("http://localhost:8000?user_id=%d", id)
	resp, err := w.Client.Get(w.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	return nil
}
