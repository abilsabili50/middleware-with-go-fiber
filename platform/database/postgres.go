package database

import (
	"fmt"
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

// create factory function
func NewDatabase(DBCfg *config.DBCfg) *Database {
	DB := createConnection(DBCfg)

	return &Database{
		Conn: DB,
	}
}

// declare create new connection function
func createConnection(DBCfg *config.DBCfg) *gorm.DB {
	// declare DB URI
	psqlcon := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Bangkok",
		DBCfg.Host,
		DBCfg.Port,
		DBCfg.User,
		DBCfg.Password,
		DBCfg.Name,
		DBCfg.SslMode,
	)

	// perform connect database with postgres driver
	db, err := gorm.Open(postgres.Open(psqlcon), &gorm.Config{})
	if err != nil {
		log.Fatalln("error occured while trying to validate database arguments:", err.Error())
	}

	log.Println("database connected")

	return db
}
