package store

import "database/sql"

type SQLiteStore interface {
	GetSQLite() *sql.DB
}
