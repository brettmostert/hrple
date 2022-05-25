package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/brettmostert/hrple/go/components/habit/internal/data"
)

func main() {
	// todo - server - http
	connString := os.Getenv("DATABASE_URL") + "/habit"

	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// this needs to be at the http server level on shutdown
	defer conn.Close()

	var activities []*data.Activity

	err = pgxscan.Select(context.Background(), conn, &activities, `SELECT id, name, type, is_archived, is_deleted, created_time, modified_time FROM activity`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(b))
}
