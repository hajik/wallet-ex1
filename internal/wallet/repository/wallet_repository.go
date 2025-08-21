package repository

import (
	"context"
	"wallet-ex1/entity"

	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *walletRepository) GetWalletByID(ctx context.Context, id int64) (*entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}
