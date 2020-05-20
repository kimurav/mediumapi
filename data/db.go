package data

import (
	"github.com/pkg/errors"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var DB sqlbuilder.Database

var settings = postgresql.ConnectionURL{
	Database: `postgres`,
	Host:     `db`,
	User:     `postgres`,
	Password: `password`,
}

func SetupDB() error {
	var err error
	DB, err = postgresql.Open(settings)
	if err != nil {
		return errors.Wrap(err, "could not connect to database")
	}
	return nil
}