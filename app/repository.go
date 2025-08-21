package app

import (
	"wallet-ex1/internal/transaction"
	trxRepository "wallet-ex1/internal/transaction/repository"
	"wallet-ex1/internal/wallet"
	walletRepository "wallet-ex1/internal/wallet/repository"

	"gorm.io/gorm"
)

type repo struct {
	walletRepository wallet.IWalletRepository
	trxRepository    transaction.ITransactionRepository
}

func (r *repo) InitRepository(db *gorm.DB) {
	r.walletRepository = walletRepository.NewWalletRepository(db)
	r.trxRepository = trxRepository.NewTransactionRepository(db)
}
