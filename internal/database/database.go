package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const DBPath = "data/hashcenter.db"

func Open() (*sql.DB, error) {

	if err := os.MkdirAll("data", 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		return nil, err
	}

	schema := `
CREATE TABLE IF NOT EXISTS dictionaries(
	path TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	size INTEGER NOT NULL,
	favorite INTEGER NOT NULL DEFAULT 0,
	priority INTEGER NOT NULL DEFAULT 1000
);
`

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
