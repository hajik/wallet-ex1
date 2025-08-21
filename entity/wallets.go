package entity

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID      int     `gorm:"primarykey;autoIncrement:true;column:id" json:"id"`
	UserID  int     `gorm:"column:user_id;not null" json:"user_id"`
	Code    string  `gorm:"uniqueIndex;column:code;not null;size:255" json:"code"`
	Balance float64 `gorm:"column:balance" json:"balance"`
}

func (r Wallet) TableName() string {
	return "public.wallets"
}
