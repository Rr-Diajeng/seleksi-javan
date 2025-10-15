package database

import (
	"fmt"
	"os"
	"seleksi-javan/model"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDBInstance() *gorm.DB {
	if db == nil {
		db = connectDB()
	}

	return db
}

func connectDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error()
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		`user=%v password=%v host=%v port=%v database=%v sslmode=disable`,
		username, password, host, port, databaseName,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()
	log.Logger = logger

	if err := DefineEnums(db); err != nil {
		log.Error().Msgf("can't define enums %s", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Task{},
	); err != nil {
		log.Error().Msgf("can't migrate tables %s", err)
	}

	return db
}

func RunMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	migrationsDir := "database/migrations/sql"

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Info().Msgf("Migrations applied successfully.")
	return nil
}
