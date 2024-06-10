package model

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(&Expense{})
	db.AutoMigrate(&User{})
}

// CreateDatabase creates a new PostgreSQL connection and returns it
func CreateDatabase() *gorm.DB {
	once.Do(func() {
		var err error
		user := viper.GetString("postgres.username")
		password := viper.GetString("postgres.password")
		dbname := viper.GetString("postgres.database")
		host := viper.GetString("postgres.host")
		port := viper.GetString("postgres.port")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		MigrateSchema(db)
	})
	return db
}
