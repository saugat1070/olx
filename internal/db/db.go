package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	db.SetMaxOpenConns(25)                 // It sets the maximum number of open connections to the database. default=0
	db.SetMaxIdleConns(25)                 // It sets the maximum number of connections in the idle connection pool. default=2
	db.SetConnMaxLifetime(5 * time.Minute) // It sets the maximum amount of time a connection may be reused.
	// fail fast
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.ping: %w", err)
	}
	return db, nil
}
