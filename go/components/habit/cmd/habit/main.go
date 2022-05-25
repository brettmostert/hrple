package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"

	"github.com/brettmostert/hrple/go/components/habit/internal/data"
)

func main() {
	// server - http
	connString := os.Getenv("DATABASE_URL") + "/habit"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
	// var name string
	// rows, err := conn.Query(context.Background(), "select name from activity")
	var activities []*data.Activity

	pgxscan.Select(context.Background(), conn, &activities, `SELECT id, name, type, is_archived, is_deleted, created_time, modified_time FROM activity`)
	// err = conn.QueryRow(context.Background(), "select name from activity").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(b))
}
