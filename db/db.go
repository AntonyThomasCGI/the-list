package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

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

func UpdateItem(id string, data map[string]interface{}) error {
	fieldValues := make([]string, 0, len(data))
	values := make([]any, 0, len(data))
	for k, v := range data {
		fieldValue := fmt.Sprintf("%s = ?", k)
		fieldValues = append(fieldValues, fieldValue)
		strVal := fmt.Sprintf("%v", v)
		values = append(values, strVal)
	}
	query := "UPDATE the_list SET " + strings.Join(fieldValues, ", ") + fmt.Sprintf(" WHERE id = %v", id)
	_, err := db.Exec(query, values...)
	if err != nil {
		logger.Error(err)
		return err
	}
	// TODO, get db.Exec res and find rows affected
	return nil
}

func SaveItem(show Show) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO the_list(title, author, watched) VALUES($1, $2, $3)",
		show.Title,
		show.Author,
		show.Watched,
	)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
