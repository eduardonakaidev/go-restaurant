package models

type User struct {
	ID            string `gorm:"type:varchar(255);primary_key" json:"id"`
	Name          string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Email         string `gorm:"type:varchar(255)" json:"email" validate:"required"`
	Password_hash string `gorm:"type:varchar(255)" json:"password" validate:"required"`
	Role          string `gorm:"type:varchar(255)" json:"role" validate:"required"`
}
