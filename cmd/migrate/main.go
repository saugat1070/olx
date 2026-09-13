package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/saugat1070/olx-api/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(("usage: migrate < up | down >"))
	}

	cfg := config.MustLoadConfig()

	m, err := migrate.New(
		"file://migration",
		cfg.DB_URL)
	if err != nil {
		log.Fatalf("migrate.new: %v", err)
	}
	switch os.Args[1] {
	case "up":
		// log.Println("migrate up called")
		if err := m.Up(); err != nil {
			log.Fatalf("migrate.up: %v", err)
		}

	case "down":
		// log.Println("migrate down called")
		if err := m.Down(); err != nil {
			log.Fatalf("migrate.down: %v", err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

	fmt.Println("Running migrations")

}
