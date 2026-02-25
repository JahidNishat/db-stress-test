package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	connStr = "postgres://postgres:12345@localhost:5432/dummy_db?sslmode=disable"
	numRows = 1000000
)

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 0. Create Table
	// file, err := os.ReadFile("schema.sql")
	// if err != nil {
	// 	log.Fatal("file reading error: ", err)
	// }
	// _, err = db.Exec(string(file))
	// if err != nil {
	// 	log.Fatal("schema exec error: ", err)
	// }
	// 1. Seed Data if needed
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if count < numRows {
		seedData(db)
	} else {
		fmt.Println("Database already seeded.")
	}

	// 2. Benchmarks
	depths := []int{100, 10000, 100000, 500000, 900000}

	fmt.Println("--- OFFSET Pagination Benchmark ---")
	for _, d := range depths {
		benchmarkOffset(db, d)
	}

	fmt.Println("--- Cursor (Keyset) Pagination Benchmark ---")
	for _, d := range depths {
		benchmarkCursor(db, d)
	}
}

func seedData(db *sql.DB) {
	fmt.Printf("Seeding %d rows... this might take a minute.", numRows)

	// Truncate first
	_, err := db.Exec("TRUNCATE users RESTART IDENTITY")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := tx.Prepare("INSERT INTO users(name, email) VALUES($1, $2)")

	for i := 1; i <= numRows; i++ {
		name := fmt.Sprintf("User%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		_, err := stmt.Exec(name, email)
		if err != nil {
			log.Fatal(err)
		}
		if i%10000 == 0 {
			fmt.Printf("Inserted %d rows...", i)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seeding complete.")
}

func benchmarkOffset(db *sql.DB, offset int) {
	start := time.Now()
	query := fmt.Sprintf("SELECT id FROM users ORDER BY id ASC LIMIT 10 OFFSET %d", offset)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
	}

	fmt.Printf("Depth: %d | Time: %v", offset, time.Since(start))
}

func benchmarkCursor(db *sql.DB, lastID int) {
	start := time.Now()
	// In a real cursor, we'd use the last ID from the previous page.
	// Here we simulate jumping to that depth.
	query := fmt.Sprintf("SELECT id FROM users WHERE id > %d ORDER BY id ASC LIMIT 10", lastID)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
	}

	fmt.Printf("Depth: %d | Time: %v", lastID, time.Since(start))
}
