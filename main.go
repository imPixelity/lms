package main

import (
	"context"
	"lms/app"
	"log"
)

func main() {
	conn, err := app.NewConn(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close()
}
