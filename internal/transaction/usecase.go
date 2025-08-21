package transaction

import (
	"context"
)

type ITransactionUsecase interface {
	Withdraw(ctx context.Context, request WithdrawRequest) error
}
