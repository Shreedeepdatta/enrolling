package initializers

import "github.com/Shreedeepdatta/rankandmarks/models"

func SyncDB() {
	DB.AutoMigrate(&models.Student{})
	DB_TEACHER.AutoMigrate(&models.Teacher{})
}
