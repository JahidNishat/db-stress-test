package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:12345@localhost:5432/dummy_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE accounts SET balance = 0, version = 1 WHERE id = 1")
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
			incrementOptimistic(db)
		}()
	}

	wg.Wait()
	log.Println("Total time taken: ", time.Since(start).Milliseconds())

	var finalBalance int
	db.QueryRow("SELECT balance FROM accounts WHERE id = 1").Scan(&finalBalance)
	fmt.Printf("Expected Balance: 100 Actual Balance: %d", finalBalance)
}

func incrementOptimistic(db *sql.DB) {
	for {
		// 1. Read current balance
		var currentBalance, version int
		err := db.QueryRow("SELECT balance, version FROM accounts WHERE id = 1").Scan(&currentBalance, &version)
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		newBalance := currentBalance + 1
		row, err := db.Exec("UPDATE accounts SET balance = $1, version = version + 1 WHERE id = 1 and version = $2", newBalance, version)
		if err != nil {
			log.Println("unknown error occured for id 1 -> version: ", version, err)
			return
		}

		r, err := row.RowsAffected()
		if err != nil {
			log.Println("rows affected error: ", err)
			return
		}

		if r == 0 {
			// log.Println("row affected 0 for version: ", version)
			continue
		}
		break
	}
}
