package itszwizard_objects

import "github.com/jinzhu/gorm"

/*
Saves a Token to an Address with the UserID in the Database.
*/
type DbForgotPassword15 struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Address string `gorm:"not null"`
	Hash    string `gorm:"not null; type:VARCHAR(500)"`
}
