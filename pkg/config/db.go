package config

import (
	"os"
	"strconv"
	"time"
)

type DBCfg struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func LoadDBCfg() *DBCfg {
	// declare db config variable
	var db = &DBCfg{}

	// read database config from environment file
	db.Host = os.Getenv("DB_HOST")
	db.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.SslMode = os.Getenv("DB_SSL_MODE")
	db.MaxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	db.MaxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	lifetime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS_MINUTES"))
	db.MaxConnLifetime = time.Duration(lifetime) * time.Minute
	idleTime, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONNECTIONS_MINUTES"))
	db.MaxConnIdleTime = time.Duration(idleTime) * time.Minute

	return db
}
