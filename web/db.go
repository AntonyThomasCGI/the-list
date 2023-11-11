package web

import (
  "os"
  "path/filepath"
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
  logger "github.com/sirupsen/logrus"
)


var db *sql.DB

type Show struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
}


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
CREATE TABLE IF NOT EXISTS the_list (id varchar(36), title varchar(255), author varchar(255));
INSERT INTO the_list(id, title, author) VALUES('1', 'lawnmower man', 'natasha');
INSERT INTO the_list(id, title, author) VALUES('2', 'paul blart mall cop', 'antony');
INSERT INTO the_list(id, title, author) VALUES('3', 'chicken run 2', 'natasha');
`
  _, err = db.Exec(sts)

  if err != nil {
    logger.Fatal(err)
  }

  logger.Info("Succesfully initialized db")
}


func getItems() ([]Show, error) {
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

