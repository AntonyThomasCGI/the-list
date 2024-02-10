package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"reflect"

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
CREATE TABLE IF NOT EXISTS the_list (id INTEGER PRIMARY KEY, title VARCHAR(255), author VARCHAR(255), watched INTEGER);
`
	_, err = db.Exec(sts)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Successfully initialized db")
}

func GetItems() ([]Show, error) {
	rows, err := db.Query("SELECT * FROM the_list")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []Show{}
	for rows.Next() {
		var show Show

		s := reflect.ValueOf(&show).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			return nil, err
		}

		result = append(result, show)
	}
	return result, nil
}

func SaveItem(show Show) (int64, error) {
	// Convert bool values to int.
	var watched int64
	if show.Watched {
		watched = 1
	} else {
		watched = 0
	}
	// Store entry.
	result, err := db.Exec("INSERT INTO the_list(title, author, watched) VALUES($1, $2, $3)", show.Title, show.Author, watched)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
