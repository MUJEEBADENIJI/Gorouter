package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

// DBConfig represents the database connection configuration
type DBConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    DBName   string
}

// NewDBConnection returns a new database connection
func NewDBConnection(config DBConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
        config.Username, config.Password, config.Host, config.Port, config.DBName)
    return sql.Open("mysql", dsn)
}

func main() {
    config := DBConfig{
        Host:     "localhost",
        Port:     3306,
        Username: "meuser",
        Password: " ",
        DBName:   "goconn",
    }

    db, err := NewDBConnection(config)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    fmt.Println("Database connection established")
}
