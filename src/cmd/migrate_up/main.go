package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// データベースのマイグレーションを行う
func main() {
	m, err := migrate.New(
		"file://db/migration",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		panic(err)
	}
	fmt.Println("done migrate")
}
