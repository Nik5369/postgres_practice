package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	UPDATE tasks 
	SET completed = TRUE
	WHERE id = 3;
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
