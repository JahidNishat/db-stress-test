package workload

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type DummyWorkload struct{}

func (w *DummyWorkload) Name() string {
	return "Dummy Workload"
}

func (w *DummyWorkload) Setup(db *sql.DB) error {
	fmt.Println("Setting up Dummy Workload...")
	return nil
}

func (w *DummyWorkload) Run(db *sql.DB) error {
	jitter := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(50*time.Millisecond + jitter)
	return nil

	// for i := 0; i < 10000; i++ {
	// 	_ = fmt.Sprintf("%d", i) // String allocation is heavy
	// }
	// return nil
}
