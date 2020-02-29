package itszwizard_objects

import "github.com/jinzhu/gorm"

type DbTest struct {
	gorm.Model
	Test1 string
	Test2 string `gorm:"unique"`
	Test3 string `gorm:"unique"`
}
