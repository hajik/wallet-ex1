package transaction

import (
	"context"
	"wallet-ex1/entity"
)

type ITransactionRepository interface {
	GetWalletByCode(ctx context.Context, code string) (*entity.Wallet, error)
	BalanceUpdate(ctx context.Context, code string, amount float64) error
	CreateTransaction(ctx context.Context, trx entity.Transaction, code string, amount float64) error
}
