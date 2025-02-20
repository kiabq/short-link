package internal

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/jackc/pgx"
)

type DatabaseMethods interface {
	CreateCode()
	CheckAlias()
	FetchAlias()
}

type DB struct {
	connection *pgx.Conn
}

func (d *DB) CreateCode(url string) error {
	var code string = ""
	chars := []string{"abcdefghijklmnopqrstuvwrxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}

	for i := 0; i < 6; i++ {
		char := int(math.Floor(rand.Float64() * float64(len(chars))))
		code += chars[char]
	}

	err := d.CheckAlias(code)
	if err != nil {
		fmt.Println("error occurred: ", err)
		return nil
	}

	fmt.Println("code: ", code)

	// do join to check if code + alias combo exists
	rows, _ := d.connection.Query(``, code)
	fmt.Println("rows: ", rows)

	return nil
}

func (d *DB) CheckAlias(code string) error {
	d.connection.Exec(`SELECT * FROM alias WHERE alias = $1`, code)

	return nil
}

func (d *DB) FetchAlias() error {
	d.connection.Query("")

	return nil
}
