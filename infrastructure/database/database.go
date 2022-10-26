package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// DbConnection create database connection
func DbConnection(masterDSN, replicaDSN string) error {
	var db = DB

	// logMode := viper.GetBool("DB_LOG_MODE")
	// debug := viper.GetBool("DEBUG")

	// loglevel := logger.Silent
	// if logMode {
	// 	loglevel = logger.Info
	// }

	// db, err = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
	// 	Logger: logger.Default.LogMode(loglevel),
	// })
	// if !debug {
	// 	db.Use(dbresolver.Register(dbresolver.Config{
	// 		Replicas: []gorm.Dialector{
	// 			postgres.Open(replicaDSN),
	// 		},
	// 		Policy: dbresolver.RandomPolicy{},
	// 	}))
	// }
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Db connection error")
		return err
	}
	DB = db
	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}
