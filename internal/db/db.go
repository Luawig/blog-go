package db

import (
	"blog-go/config"
	"blog-go/internal/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbConfig := config.GetDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema, this will create table if they don't exist
	_ = DB.AutoMigrate(&model.Article{}, &model.Category{}, &model.Comment{}, &model.User{})

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

func InitTestDB() {
	var err error

	dbConfig := config.GetDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database+"_test",
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = DB.Migrator().DropTable(&model.Article{}, &model.Category{}, &model.Comment{}, &model.User{})
	// Migrate the schema, this will create table if they don't exist
	_ = DB.AutoMigrate(&model.Article{}, &model.Category{}, &model.Comment{}, &model.User{})

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
