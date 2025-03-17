package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(connString string) error {
	var err error
	DB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	return nil
}
