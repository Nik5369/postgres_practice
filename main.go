package main

import (
	"context"
	"postgres_connection/feature_postgres/simple_connection"
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

	// if err := simple_sql.CreateTable(conn, ctx); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Table created successfully")
	// }

	// if err := simple_sql.InsertRow(
	// 	conn,
	// 	ctx,
	// 	simple_sql.TaskModel{
	// 		Title:       "погулять с кошкой",
	// 		Description: "надо погулять с кошкой 3 раза за день",
	// 		Completed:   false,
	// 		Created_at:  time.Now()},
	// ); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Insert successfully")
	// }

	// tasks, err := simple_sql.SelectRows(conn, ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// task := tasks[4]
	// task.Completed = true
	// if err := simple_sql.UpdateTask(conn, ctx, task); err != nil {
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
