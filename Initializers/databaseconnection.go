package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_TEACHER *gorm.DB

func DatabaseConn() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
}
func TeacherDatabaseConn() {
	var err error
	dsn := os.Getenv("DB_TEACHER")
	DB_TEACHER, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}
}
