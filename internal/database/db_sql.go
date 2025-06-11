package database

import "database/sql"

func OpenDbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "customers.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		phone TEXT,
		address TEXT,
		listing_link TEXT,
		notes TEXT,
		type TEXT
	)`)

	if err != nil {
		return nil, err
	}
	return db, nil
}
