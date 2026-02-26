package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp"
)

func SelectRows(conn *pgx.Conn, ctx context.Context) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC
	;`

	rows, err := conn.Query(ctx, sqlQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		pp.Println(task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
