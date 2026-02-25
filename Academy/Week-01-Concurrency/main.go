package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for range ch {
		count++
	}

	fmt.Println("Total count:", count)

}
