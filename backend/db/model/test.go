package model

import "github.com/jinzhu/gorm"

// Test test struct for db connection
type Test struct {
	gorm.Model
	Message string `gorm:"column:test_msg" json:"message"`
}
