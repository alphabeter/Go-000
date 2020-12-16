package model

import "gorm.io/gorm"

// User struct.
type User struct {
	gorm.Model
	Name string
	Age  uint
}
