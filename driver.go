package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Declaring some connection pool properties
const (
	maxOpenDbConns = 10
	maxIdleDbConns = 5
	maxDbLifeTime  = 5 * time.Minute
)

// ConnectSQL returns a connection pull with customized settings
func ConnectSQL(dsn string) (*sql.DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDbConns)
	d.SetMaxIdleConns(maxIdleDbConns)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// testDB pings database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return fmt.Errorf("error: %v while testing a connection", err)
	}
	return nil
}

// NewDataBase creates a new connection pull  with pgx driver
func NewDataBase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error: %v while connecting to database", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error: %v while pinging database", err)

	}
	return db, nil
}
