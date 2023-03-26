package repo

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	client *sqlx.DB
}

func NewRepo() (*Repo, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"user",
		"password",
		"host",
		"port",
		"DBName",
	)

	log.Println("connecting to db...")

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	log.Println("connected to db")

	return &Repo{db}, nil
}
