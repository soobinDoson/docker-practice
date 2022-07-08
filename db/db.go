package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	engine *gorm.DB
}

// ConnectDb open connection to db
// ConnectDb expose ...
func (d *DB) ConnectDb(sqlDSN string) error {
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  sqlDSN,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			// NamingStrategy: schema.NamingStrategy{
			// 	TablePrefix:   "partner.", // schema name
			// 	SingularTable: false,
			// },
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			if err := sqlDb.Ping(); err != nil {
				log.Print(err)
			}
		}
	}()
	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDb.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDb.SetMaxOpenConns(100)
	d.engine = db
	return nil
}
