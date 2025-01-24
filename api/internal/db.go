package internal

import "github.com/jackc/pgx"

type DatabaseMethods interface {
	CreateCode()
	CheckAlias()
	FetchAlias()
}

type DB struct {
	connection *pgx.Conn
}

func (d *DB) CreateCode(url string) {
	var code string = ""
	chars := []string{"abcdefghijklmnopqrstuvwrxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}

	// do join to check if code + alias combo exists
	d.connection.Query(`SELECT * FROM alias WHERE alias = $1`, code)
}

func (d *DB) CheckAlias(code string) {
	d.connection.Exec(`SELECT * FROM alias WHERE alias = $1`, code)
}

func (d *DB) FetchAlias() {
	d.connection.Query("")
}
