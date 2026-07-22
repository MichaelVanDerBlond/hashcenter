package database

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const DBPath = "data/hashcenter.db"

const schema = `
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;
PRAGMA synchronous = NORMAL;

CREATE TABLE IF NOT EXISTS dictionaries(
	path TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	size INTEGER NOT NULL,
	favorite INTEGER NOT NULL DEFAULT 0,
	priority INTEGER NOT NULL DEFAULT 1000
);

CREATE TABLE IF NOT EXISTS sessions(
	id TEXT PRIMARY KEY,
	state TEXT NOT NULL,

	started TEXT,
	finished TEXT,

	interface TEXT,
	channel INTEGER,

	backend TEXT,

	capture_file TEXT,

	error TEXT
);
`

func Open() (*sql.DB, error) {

	if err := os.MkdirAll("data", 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", DBPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
