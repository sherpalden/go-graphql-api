package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env, zapLogger Logger) Database {
	username := env.DBUser
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		zapLogger.Zap.Info("Url: ", dsn)
		zapLogger.Zap.Panic(err)
	}

	zapLogger.Zap.Info("Database connection established")

	return Database{
		DB: db,
	}
}
