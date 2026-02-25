package main

import (
	"context"
	"fmt"
	"postgres_connection/feature_postgres/simple_connection"
	"postgres_connection/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()

	conn, pgxErr := simple_connection.InitConnection(ctx)
	if pgxErr != nil {
		panic(pgxErr)
	}

	// if err := conn.Ping(ctx); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.CreateTable(conn, ctx); err != nil {
		panic(err)
	} else {
		fmt.Println("Table created successfully")
	}

	// if err := simple_sql.InsertRow(conn, ctx, "обед", "покушал", false, time.Now()); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Insert successfully")
	// }

	// if err := simple_sql.UpdateRow(conn, ctx); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Update successfully")
	// }

	// if err := simple_sql.DeleteRow(conn, ctx); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Delete successfully")
	// }
}
