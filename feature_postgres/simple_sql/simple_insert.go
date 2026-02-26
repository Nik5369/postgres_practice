package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(conn *pgx.Conn, ctx context.Context, tasks TaskModel) error {
	sqlQuery := `
	INSERT INTO tasks 
	( title, description, completed, created_at )
	 VALUES ($1, $2, $3, $4);`

	_, err := conn.Exec(ctx, sqlQuery, tasks.Title, tasks.Description, tasks.Completed, tasks.Created_at)
	return err
}
