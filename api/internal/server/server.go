package server

import (
	"fmt"
	"kiabq/hyperslice/internal/routes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load(".env.local")
	if err != nil {
		// TODO: Do a graceful shutdown if this occurs
		fmt.Errorf("Error loading environment variables: %w", err)
	}

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 16)
	if err != nil {
		// TODO: Do a graceful shutdown if this occurs
		fmt.Errorf("Error parsing port to 16 bit uint: %w", err)
	}

	config := pgx.ConnConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     uint16(port),
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Errorf("Error connecting to Postgres database: %+w\n", err)
	}

	// Do something with this connection eventually
	fmt.Println(conn)

	// start server
	routes.RegisterRoutes()
	log.Fatal(s.ListenAndServe())
}
