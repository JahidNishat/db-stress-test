package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

const (
	connStr = "postgres://postgres:12345@localhost:5432/dummy_db?sslmode=disable"
)

func transfer(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("txn begin error: %w", err)
	}
	defer tx.Rollback()

	// Setting isolation level repeatable read
	_, err = tx.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;")
	if err != nil {
		return fmt.Errorf("setting isolation level repeatable read error: %w", err)
	}

	// Read alice and bob balance
	var aliceBal, bobBal int
	err = tx.QueryRow("SELECT balance FROM accounts where id = 1").Scan(&aliceBal)
	if err != nil {
		return fmt.Errorf("read alice balance error: %w", err)
	}

	err = tx.QueryRow("SELECT balance FROM accounts where id = 2").Scan(&bobBal)
	if err != nil {
		return fmt.Errorf("read bob balance error: %w", err)
	}

	time.Sleep(2 * time.Second)

	// Update alice and bob new balance
	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE id = 1", aliceBal-50)
	if err != nil {
		return fmt.Errorf("alice balance update error: %w", err)
	}

	_, err = tx.Exec("UPDATE accounts SET balance = $1 WHERE id = 2", bobBal+50)
	if err != nil {
		return fmt.Errorf("bob balance update error: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit error: %w", err)
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening db: ", err)
	}
	defer db.Close()

	log.Println("---- Starting Transaction ----")

	for {
		err = transfer(db)
		if err == nil {
			log.Println("---- Success ----")
			return
		}

		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "40001" {
			log.Println("*** Conflict Detected... Retrying... ***")
			continue
		}
		log.Fatal("Unrecoverable error: ", err)
	}
}
