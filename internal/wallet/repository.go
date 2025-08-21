package wallet

import (
	"context"
	"wallet-ex1/entity"
)

type IWalletRepository interface {
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	GetWalletByID(ctx context.Context, id int64) (*entity.Wallet, error)
}
