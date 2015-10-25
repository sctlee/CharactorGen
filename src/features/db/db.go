package db

import (
	"log"

	"github.com/jackc/pgx"
)

var Pool *pgx.ConnPool

func StartPool() {
	var err error
	Pool, err = pgx.NewConnPool(extractConfig())
	if err != nil {
		log.Fatalln("Unable to connect to database")
	}
}

func extractConfig() pgx.ConnPoolConfig {
	var config pgx.ConnPoolConfig

	config.ConnConfig = *noPasswordConnConfig

	return config
}
