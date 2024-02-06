package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	logger "github.com/sirupsen/logrus"
)

var db *sql.DB

func init() {
	dbpath := "/etc/the-list"
	err := os.MkdirAll(dbpath, 0777)
	if err != nil {
		logger.Fatal("Failed to create db directory: ", err)
	}

	db, err = sql.Open("sqlite3", filepath.Join(dbpath, "the-list.db"))
	if err != nil {
		logger.Fatal(err)
	}

	// Create table if it does not exist.
	sts := `
CREATE TABLE IF NOT EXISTS the_list (id INTEGER PRIMARY KEY, title varchar(255), author varchar(255));
`
	_, err = db.Exec(sts)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Succesfully initialized db")
}

func GetItems() ([]Show, error) {
	rows, err := db.Query("SELECT * FROM the_list")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []Show{}
	for rows.Next() {
		var row Show

		err = rows.Scan(&row.ID, &row.Title, &row.Author)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func SaveItem(show Show) (int64, error) {
	result, err := db.Exec("INSERT INTO the_list(title, author) VALUES($1, $2)", show.Title, show.Author)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
