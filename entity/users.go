package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primarykey;autoIncrement:true;column:id" json:"id"`
	Email    string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Username string `gorm:"column:username;uniqueIndex;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
}

func (r User) TableName() string {
	return "public.users"
}
