package starter

import (
	"github.com/IMQS/log"
	"github.com/IMQS/nf/nfdb"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Keep a copy of your schema at https://dbdiagram.io. For an example, see https://dbdiagram.io/d/5d9f86ffff5115114db5209a

var migrationsSQL []string

func openDB(log *log.Logger, dbConf nfdb.DBConfig) (*gorm.DB, error) {
	var flags nfdb.DBConnectFlags
	flags |= nfdb.DBConnectFlagWipeDB
	db, err := nfdb.OpenDB(log, dbConf.Driver, dbConf.DSN(), nfdb.MakeMigrations(log, migrationsSQL), flags)
	if err != nil {
		return nil, err
	}
	// If you want to avoid using migrations for early prototyping, and instead just want GORM
	// to generate your DB schema for you, then use db.AutoMigrate(), like below.
	db.AutoMigrate(&frog{}, &frogType{})
	return db, nil
}

func init() {
	migrationsSQL = []string{
		`
		-- We're just going to use GORM AutoMigrate until we have nailed down our DB schema more,
		-- at which point we switch to using more formal migrations, like this:
		
		-- CREATE EXTENSION IF NOT EXISTS postgis;
		-- CREATE TABLE frog (id BIGSERIAL PRIMARY KEY, description VARCHAR, frog_type_id BIGINT NOT NULL);
		-- CREATE TABLE frog_type (id BIGSERIAL PRIMARY KEY, description VARCHAR);
		-- CREATE UNIQUE INDEX idx_frog_type_description ON frog_type (description);
		`,
	}
}
