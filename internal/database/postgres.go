package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

func NewConnection() (*DB, error) {
    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "localhost"
    }
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "5432"
    }
    user := os.Getenv("DB_USER")
    if user == "" {
        user = "postgres"
    }
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    if dbname == "" {
        dbname = "kongpay"
    }
    sslmode := os.Getenv("DB_SSL_MODE")
    if sslmode == "" {
        sslmode = "disable"
    }

    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode,
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("✅ PostgreSQL connected")
    return &DB{db}, nil
}

func (db *DB) Close() error {
    return db.DB.Close()
}
