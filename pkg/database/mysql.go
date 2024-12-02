package database

import (
	"fmt"
	"log"
	"os"
	"rentalMobil/internal/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToMysql() *gorm.DB {
	dbUSER := os.Getenv("DB_USER")
	dbPASWORD := os.Getenv("DB_PASWORD")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbDBNAME := os.Getenv("DB_DBNAME")

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUSER, dbPASWORD, dbHOST, dbPORT, dbDBNAME)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // Output to standard logger
		logger.Config{
			SlowThreshold: time.Second, // Log queries slower than this
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color output
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil
	}

	db.AutoMigrate(&models.Users{}, &models.Cars{}, &models.Booking{})

	return db
}
