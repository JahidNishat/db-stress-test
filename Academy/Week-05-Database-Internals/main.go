package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const connStr = "postgres://postgres:12345@localhost:5432/dummy_db?sslmode=disable"

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Reset Balance
	_, err = db.Exec("UPDATE accounts SET balance = 0 WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting 100 workers to increment balance...")

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementBalance(db)
		}()
	}

	wg.Wait()
	log.Println("Total time taken: ", time.Since(start).Milliseconds())

	var finalBalance int
	db.QueryRow("SELECT balance FROM accounts WHERE id = 1").Scan(&finalBalance)
	fmt.Printf("Expected Balance: 100 Actual Balance: %d", finalBalance)
}

func incrementBalance(db *sql.DB) {
	// 1. Read current balance
	var currentBalance int
	tx, err := db.Begin()
	if err != nil {
		log.Println("Transaction Begin Error: ", err)
		return
	}
	defer tx.Rollback()

	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = 1 FOR UPDATE").Scan(&currentBalance)
	if err != nil {
		log.Println("Read error:", err)
		return
	}

	// 2. Simulate some local processing time
	// time.Sleep(10 * time.Millisecond)

	// 3. Write back the incremented balance
	newBalance := currentBalance + 1
	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE id = 1", newBalance)
	if err != nil {
		log.Println("Write error:", err)
	}
	tx.Commit()
}
