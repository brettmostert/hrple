package common

import "github.com/jackc/pgx/v4/pgxpool"

type AppContext struct {
	Db *pgxpool.Pool
}
