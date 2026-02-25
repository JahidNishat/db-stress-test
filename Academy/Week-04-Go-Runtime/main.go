package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// A CPU-bound task
func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Just do some math to keep the CPU busy
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	_ = sum
}

func main() {
	// 1. Create a trace file
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 2. Start tracing
	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Starting 1000 CPU-bound goroutines...")

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done! Now run: go tool trace trace.out")
}
