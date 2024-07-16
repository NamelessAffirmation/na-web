package database

import (
	"context"
	"fmt"
)

func CreateStagedMessage(message string) {
	_, err := DB.Exec(context.Background(), "INSERT INTO public.staged_messages (message) VALUES ($1)", message)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Print("ERROR")
	}
}

// func GetStagedMessage() {
// 	var name string
// 	var weight int64
// 	_, err = DB.Exec(context.Background(), "")
// 	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
// }
