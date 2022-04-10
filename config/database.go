package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {

	// Load Environment
	loadEnvironmentSetting()

	// Setting Env Database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	levelDev := os.Getenv("LEVEL_DEV")

	var dsn string
	if levelDev == "development" {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta`, dbHost, dbUser, dbPass, dbPort, dbName)
	} else {
		dsn = fmt.Sprintf(`host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Jakarta`, dbHost, dbUser, dbPass, dbPort, dbName)
	}

	// Open DB Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB Connection Failed")
	} else {
		fmt.Println("DB Connection Success")
	}

	//migrateDatabase(db)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	database.Close()
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate()
}

func loadEnvironmentSetting() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error read .env file")
	} else {
		fmt.Println("Success read .env file")
	}
}
