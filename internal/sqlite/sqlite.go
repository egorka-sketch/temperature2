package sqlite

import "database/sql"

type storage struct {
	db *sql.DB
}
