package session

import "database/sql"

type Session struct {
	db *sql.DB
}
