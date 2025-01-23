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
)

func Start() {
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 16)

	if err != nil {
		// throw error
	}

	config := pgx.ConnConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     uint16(port),
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
	}

	conn, err := pgx.Connect(config)
	fmt.Println(conn)

	routes.RegisterRoutes()

	log.Fatal(s.ListenAndServe())
}
