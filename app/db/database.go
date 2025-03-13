package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func loadEnv() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
    config.SetConfigType("env")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

    return config
}

func NewDB() *sql.DB {
    config := loadEnv()

	host := config.GetString("DB_HOST")
	port := config.GetInt("DB_PORT")
	username := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	dbName := config.GetString("DB_NAME")

	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatalf("Connected database error, %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
