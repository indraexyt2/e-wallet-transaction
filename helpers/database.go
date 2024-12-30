package helpers

import (
	"e-wallet-transaction/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func SetupMySql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_HOST", "localhost"), GetEnv("DB_PORT", "3306"), GetEnv("DB_NAME", "e-wallet"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	logrus.Info("Database initiated using gorm")

	err = DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	logrus.Info("Database migrated")
}
