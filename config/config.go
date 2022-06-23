package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func NewPostgresPool(host, port, user, password, name string) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	return pgxpool.Connect(context.Background(), connCfg)
}

func Config() (*pgxpool.Pool, error) {
	pg, err := NewPostgresPool(os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to database")
	
	return pg, nil
}
