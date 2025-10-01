package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"time"
)

func ConnectionDatabase(viper *viper.Viper, logs *zap.Logger) *sqlx.DB {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logs.Fatal(err.Error())
	}

	// setup connection pool

	db.SetMaxIdleConns(10)                  // idle connections
	db.SetMaxOpenConns(100)                 // max open connections
	db.SetConnMaxLifetime(30 * time.Minute) // lifetime tiap connection
	db.SetConnMaxIdleTime(5 * time.Minute)  // idle max time

	// test connection
	if err := db.Ping(); err != nil {
		logs.Fatal("failed to ping DB", zap.Error(err))
	}
	logs.Info("Database connected with pool setup")

	return db
}
