package itszwizard_objects

import "github.com/jinzhu/gorm"

/*
In this struct the SessionID and IPAdress ist stored also the create, delete and update time
*/
type DbSession15 struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	SessionId string `gorm:"not null;type:VARCHAR(50)"`
	IPAddress string `gorm:"not null; type:VARCHAR(500)"`
}
