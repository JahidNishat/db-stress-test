package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	tokens     float64
	maxTokens  float64
	refillRate float64
	lastRefill time.Time
}

func NewRateLimiter(maxTokens, refillRate float64) *RateLimiter {
	return &RateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elaped := now.Sub(rl.lastRefill)
	refillTokens := elaped.Seconds() * rl.refillRate
	if refillTokens > 0 {
		rl.tokens += refillTokens
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastRefill = now
	}

	if rl.tokens >= 1.0 {
		rl.tokens -= 1.0
		return true
	}
	return false
}

func main() {
	// burst 10 requests, refill 5 tokens per second
	limiter := NewRateLimiter(10, 5)
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if limiter.Allow() {
				println("Request", i, "allowed")
			} else {
				println("Request", i, "denied")
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
