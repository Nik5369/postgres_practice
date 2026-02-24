package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//"postgres://YourUserName:YourPassword@YouHostName:5432/YouDatabaseName"


func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:d310301@localhost:5432/postgres")	
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connection success")
}
