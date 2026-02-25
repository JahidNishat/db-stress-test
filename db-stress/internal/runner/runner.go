package runner

import (
	"context"
	"database/sql"
	"db-stress/internal/workload"
	"fmt"
	"sync"
	"time"
)

var (
	successCount int64
	failCount    int64
)

type Stats struct {
	Duration time.Duration
	IsError  bool
}

type Config struct {
	Workers  int
	Duration time.Duration
	DSN      string
}

func Run(cfg Config, wl workload.Workload, statsCh chan<- Stats) error {
	// fmt.Printf("🚀 initializing runner with %d worker for %v...\n", cfg.Workers, cfg.Duration)

	var db *sql.DB
	if cfg.DSN != "" {
		var err error
		db, err = sql.Open("postgres", cfg.DSN)
		if err != nil {
			return fmt.Errorf("db connection failed: %w", err)
		}
		defer db.Close()
	} else {
		// fmt.Println("⚠️ NO DSN provided. Running in simulation mode")
	}

	if err := wl.Setup(db); err != nil {
		return fmt.Errorf("workload setup failed: %w", err)
	}

	// fmt.Println("Starting Worker... ... ...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Duration)
	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= cfg.Workers; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					start := time.Now()
					err := wl.Run(db)
					took := time.Since(start)

					statsCh <- Stats{
						Duration: took,
						IsError:  err != nil,
					}
				}
			}
		}(i)
	}

	wg.Wait()
	close(statsCh)
	// fmt.Println("Stress test completed... ... ...")
	return nil
}
