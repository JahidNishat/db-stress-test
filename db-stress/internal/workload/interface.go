package workload

import (
	"database/sql"
)

type Workload interface {
	Name() string
	Setup(db *sql.DB) error
	Run(db *sql.DB) error
}
