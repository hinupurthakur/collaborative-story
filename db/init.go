package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// db connection variable to communicate to database
var db *sqlx.DB

func InitDB() {
	log.Infoln("initializing db")
	var err error
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	psqlDSN := os.Getenv("POSTGRES_DSN")
	db, err = sqlx.Open("postgres", psqlDSN)
	if err != nil {
		log.Fatal("unable to connect to the database, ", psqlDSN, err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("unable to ping db", err)
	}
	log.Infoln("db successfully connected!")
	MigrationsUp(db)
}
