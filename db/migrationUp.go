package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
)

func MigrationsUp(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalln("unable to fetch migrations", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres", driver)
	if err != nil {
		log.Fatalln("unable to fetch migrations", err)
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("unable to run migrations", err)
		return
	}
	log.Infoln("migration up successful")
}
