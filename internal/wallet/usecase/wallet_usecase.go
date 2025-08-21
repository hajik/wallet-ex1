package usecase

import (
	"context"
	"fmt"
	"wallet-ex1/internal/wallet"

	"go.uber.org/zap"
)

type walletUsecase struct {
	repo wallet.IWalletRepository
}

func NewWalletUsecase(repo wallet.IWalletRepository) *walletUsecase {
	return &walletUsecase{repo: repo}
}

func (u *walletUsecase) GetWallet(ctx context.Context, userID int64) (wallet.GetWalletResponse, error) {
	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		fmt.Errorf("failed to get user by id", zap.Error(err))
		return wallet.GetWalletResponse{}, err
	}

	balance, err := u.repo.GetWalletByID(ctx, int64(user.ID))
	if err != nil {
		fmt.Errorf("failed to get wallet by id", zap.Error(err))
		return wallet.GetWalletResponse{}, err
	}

	result := wallet.GetWalletResponse{
		Username: user.Username,
		Code:     balance.Code,
		Balance:  float64(balance.Balance),
	}

	return result, nil
}
