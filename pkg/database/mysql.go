package database

import (
	"fmt"
	"os"
	"rentalMobil/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMysql() *gorm.DB {
	dbUSER := os.Getenv("DB_USER")
	dbPASWORD := os.Getenv("DB_PASWORD")
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbDBNAME := os.Getenv("DB_DBNAME")

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUSER, dbPASWORD, dbHOST, dbPORT, dbDBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return nil
	}

	db.AutoMigrate(&models.Users{}, &models.Cars{})

	return db
}
