package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. Read file
	content, err := os.ReadFile("script.lua")
	if err != nil {
		log.Fatal("script reading error: ", err)
	}
	luaScript := string(content)

	// 2. Connect Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// 3. Config
	key := "user:test_1"
	capacity := 10 // Max 10 tokens
	rate := 1      // 1 token per second

	log.Println("Starting burst of 15 requests")
	for i := 1; i <= 15; i++ {
		now := time.Now().UnixMicro()

		val, err := rdb.Eval(ctx, luaScript, []string{key}, rate, capacity, now).Result()
		if err != nil {
			log.Fatal("redis eval error: ", err)
		}

		if val.(int64) == 1 {
			log.Printf("Req %d: Allowed\n", i)
		} else {
			log.Printf("Req %d: Denied\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
