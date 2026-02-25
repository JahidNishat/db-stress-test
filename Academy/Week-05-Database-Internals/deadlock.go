package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func transferAtoB(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("txn begin error: ", err)
		return
	}
	defer tx.Rollback()

	// 1. Lock Alice
	log.Println("[AtoB] Locking Alice (ID 1)...")
	var balA int
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = 1 FOR UPDATE").Scan(&balA)
	if err != nil {
		log.Println("[AtoB] Alice Read error:", err)
		return
	}

	time.Sleep(1 * time.Second) // Wait for B to lock Bob

	// 2. Try to Lock Bob
	log.Println("[AtoB] Trying to lock Bob (ID 2)...")
	var balB int
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = 2 FOR UPDATE").Scan(&balB)
	if err != nil {
		log.Println("[AtoB] Bob Read error (EXPECTED DEADLOCK):", err)
		return
	}

	tx.Commit()
	log.Println("[AtoB] Success")
}

func transferBtoA(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("txn begin error: ", err)
		return
	}
	defer tx.Rollback()

	// 1. Lock Bob
	log.Println("[BtoA] Locking Bob (ID 2)...")
	var balB int
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = 2 FOR UPDATE").Scan(&balB)
	if err != nil {
		log.Println("[BtoA] Bob Read error:", err)
		return
	}

	time.Sleep(1 * time.Second) // Wait for A to lock Alice

	// 2. Try to Lock Alice
	log.Println("[BtoA] Trying to lock Alice (ID 1)...")
	var balA int
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = 1 FOR UPDATE").Scan(&balA)
	if err != nil {
		log.Println("[BtoA] Alice Read error (EXPECTED DEADLOCK):", err)
		return
	}

	tx.Commit()
	log.Println("[BtoA] Success")
}

func main() {
	connStr := "postgres://postgres:12345@localhost:5432/dummy_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("--- Starting Deadlock Simulation ---")
	go transferAtoB(db)
	go transferBtoA(db)

	time.Sleep(10 * time.Second) // Deadlock detection takes a few seconds
	log.Println("Finish")
}