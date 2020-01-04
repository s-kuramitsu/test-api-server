package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique; not null" binding:"required,max=200,email"`
	Name  string `json:"name" gorm:"not null" binding:"required,max=100"`
}

type Area struct {
	gorm.Model
	Name string `json:"name" gorm:"not null" binding:"required,max=100"`
}

type CheckList struct {
	gorm.Model
	AreaId      uint      `json:"area_id" binding:"required"`
	CheckTime   time.Time `json:"check_time" gorm:"not null"`
	DoorLock    bool      `json:"door_lock"`
	Fire        bool      `json:"fire"`
	DevicePower bool      `json:"device_power"`
}
