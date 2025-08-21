package repository

import (
	"context"
	"wallet-ex1/entity"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetWalletByCode(ctx context.Context, code string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *transactionRepository) BalanceUpdate(ctx context.Context, code string, amount float64) error {
	var wallet entity.Wallet
	return r.db.WithContext(ctx).Model(&wallet).Where("code = ?", code).Update("balance", amount).Error
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, trx entity.Transaction, code string, amount float64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := trx
		if err := tx.Create(&user).Error; err != nil {
			// If there's an error, return it. GORM will automatically rollback.
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("code = ?", code).Update("balance", amount).Error; err != nil {
			// If there's an error here, GORM will rollback both the user and the wallet creation.
			user.Status = "failed"
			if err := tx.Save(&user).Error; err != nil {
				// If there's an error here, GORM will rollback both the user and the wallet creation.
				return err
			}
			return err
		}

		user.Status = "success"
		if err := tx.Save(&user).Error; err != nil {
			// If there's an error here, GORM will rollback both the user and the wallet creation.
			return err
		}

		return nil
	})
}
