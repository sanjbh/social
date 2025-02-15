package main

import (
	"fmt"
	"log"

	"github.com/sanjbh/social/internal/db"
	"github.com/sanjbh/social/internal/env"
	"github.com/sanjbh/social/internal/store"
)

const VERSION string = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetEnvVar("ADDR", ":3000"),
		db: dbConfig{
			addr:         env.GetEnvVar("DB_ADDR", "postgress://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetEnvVar("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvVar("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnvVar("DB_MAX_IDLE_TIME", "15min"),
		},
		env: env.GetEnvVar("ENV", "development"),
	}

	fmt.Println(cfg.db.addr)

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewPostgresStorage(db)

	application := &application{
		config: cfg,
		store:  store,
	}

	mux := application.mount()

	log.Fatal(application.run(mux))

}
