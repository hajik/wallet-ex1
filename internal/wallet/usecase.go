package wallet

import (
	"context"
)

type IWalletUsecase interface {
	GetWallet(ctx context.Context, userID int64) (GetWalletResponse, error)
}
