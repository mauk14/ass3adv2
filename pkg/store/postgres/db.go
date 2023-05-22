package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDb(host, port, user, password, dbname string) (pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s")
}
