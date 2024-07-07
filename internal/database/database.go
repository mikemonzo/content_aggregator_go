package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	// Database connection
	connection **sql.DB
}

func NewDatabase(dbName string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	database := &Database{connection: &db}
	return database, database.createTable()
}

func (d *Database) createTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS content (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        
        "data" TEXT
    );`
	_, err := (*d.connection).Exec(createTableSQL)
	return err
}

func (d *Database) InsertContent(data string) error {
	insertSQL := `INSERT INTO content(data) VALUES(?)`
	_, err := (*d.connection).Exec(insertSQL, data)
	return err
}

func (d *Database) FetchAllContent() ([]string, error) {
	var data string
	var content []string
	rows, err := (*d.connection).Query("SELECT data FROM content")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		content = append(content, data)
	}
	return content, nil
}

func (d *Database) Close() error {
	return (*d.connection).Close()
}
