package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "sql11705892:9CyEUmc1yS@tcp(sql11.freemysqlhosting.net:3306)/sql11705892")
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to the database")
    DB = db
    return DB, nil
}
