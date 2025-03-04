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

   err = CreateTables(db)
   if err != nil {
       panic(err)
   }

   fmt.Println("Database tables created")
}

// CreateTables creates the device and activity log tables in the database
func CreateTables(db *sql.DB) error {
    // Create devices table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS devices (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            mac_address VARCHAR(255) NOT NULL,
            ip_address VARCHAR(255) NOT NULL
        );
    `)
    if err != nil {
        return err
    }

    // Create activity logs table
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS activity_logs (
            id INT AUTO_INCREMENT PRIMARY KEY,
            device_id INT NOT NULL,
            visited_url VARCHAR(255) NOT NULL,
            timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (device_id) REFERENCES devices(id)
        );
    `)
    return err
}
