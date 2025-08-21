package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID              int       `gorm:"primarykey;autoIncrement:true;column:id" json:"id"`
	WalletID        int       `gorm:"column:wallet_id;not null" json:"wallet_id"`
	Type            string    `gorm:"column:type;not null;size:255" json:"type"`
	Amount          float64   `gorm:"column:amount" json:"amount"`
	Status          string    `gorm:"column:status" json:"status"`
	TransactionDate time.Time `gorm:"column:transaction_date" json:"transaction_date"`
}

func (r Transaction) TableName() string {
	return "public.transactions"
}
