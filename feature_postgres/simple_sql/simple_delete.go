package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = 5
	;`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
