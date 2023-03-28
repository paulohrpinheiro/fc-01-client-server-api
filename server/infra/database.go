package infra

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./exchange.sqlite3")
	if err != nil {
		return nil, fmt.Errorf("error on sql.Open: %v", err)
	}

	return db, nil
}

func Write(bid string) error {
	db, err := connect()
	if err != nil {
		return fmt.Errorf("sql connect error: %v", err)
	}

	defer db.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	_, err = db.ExecContext(ctx, "INSERT INTO exchange(bid) VALUES(?)", bid)
	if err != nil {
		return fmt.Errorf("sql insert error: %v", err)
	}

	return nil
}
