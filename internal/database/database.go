package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/vss414/hw-1/internal/config"
	"log"
)

func New() *pgxpool.Pool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	psqlConn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbName,
	)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("failed to ping database", err)
	}

	return pool
}
