package db

import (
	"coursehub/models"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB
var err error

func ConstructDSN(filename string) (string, error) {
	var DatabaseConfig Config
	if _, err := toml.DecodeFile(filename, &DatabaseConfig); err != nil {
		panic(err)
	}
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DatabaseConfig.Username,
		DatabaseConfig.Password, DatabaseConfig.Address, DatabaseConfig.Port, DatabaseConfig.Database)
	return DSN, err
}

func Init() {
	filename := os.Getenv("PATH_TO_CONFIG") // set up the environment variable required
	DSN, _ := ConstructDSN(filename)
	db, _ = gorm.Open("mysql", DSN)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(
		&models.CourseModel{},
		&models.GroupModel{},
	)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
