package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YouHostName:5432/YouDatabaseName"

func InitConnection(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://postgres:d310301@localhost:5432/postgres")

}
