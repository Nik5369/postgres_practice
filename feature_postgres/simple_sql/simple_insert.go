package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(conn *pgx.Conn, ctx context.Context, title string, description string, completed bool, createdAt time.Time) error {
	sqlQuery := `
	INSERT INTO tasks 
	( title, description, completed, created_at )
	 VALUES ($1, $2, $3, $4);`

	_, err := conn.Exec(ctx, sqlQuery, title, description, completed, createdAt)
	return err
}
