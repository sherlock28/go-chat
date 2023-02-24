package db

import (
	"database/sql"
	"fmt"
	"log"
	"server/config"
	"strconv"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(config *config.Configuration) (*Database, error) {
	log.Println("⌛ Initializing database connection...")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DB.User, config.DB.Password, config.DB.Host, strconv.Itoa(config.DB.Port), config.DB.Name, config.DB.SSLMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("❌ Could not initialize database connection: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Could not initialize database connection: %s", err)
	}

	log.Println("✔️ Database connected")

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
