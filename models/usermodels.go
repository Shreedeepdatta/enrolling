package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string
	Roll     int16 `gorm:unique`
	Class    string
	Password string
}

type Teacher struct {
	gorm.Model
	Name           string
	Subject        string
	Qualifications string
	Experience     string
	Password       string
}
