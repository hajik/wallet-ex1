package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wallet-ex1/entity"
	"wallet-ex1/internal/transaction"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type transactionUsecase struct {
	repo transaction.ITransactionRepository
}

func NewTransactionUsecase(repo transaction.ITransactionRepository) *transactionUsecase {
	return &transactionUsecase{repo: repo}
}

func (u *transactionUsecase) Withdraw(ctx context.Context, request transaction.WithdrawRequest) error {

	wallet, err := u.repo.GetWalletByCode(ctx, request.Code)
	if err != nil {
		log.Error("failed to get wallet by code. ", err)
		return errors.New("Wallet not found")
	}

	var minAmount float64 = viper.GetFloat64("transaction.withdraw.min")

	// Check minimum amount
	if request.Amount < minAmount {
		log.Error("Failed amount higher then balance.")
		return errors.New(fmt.Sprintf("Minimum amount Rp %v", minAmount))
	}

	// Check balance amount
	if request.Amount > wallet.Balance {
		log.Error("Failed amount higher then balance.")
		return errors.New("Your balance is insufficient")
	}

	balanceReduction := wallet.Balance - request.Amount

	trx := entity.Transaction{
		WalletID:        wallet.ID,
		Type:            "withdraw",
		Amount:          request.Amount,
		Status:          "pending",
		TransactionDate: time.Now(),
	}

	if err := u.repo.CreateTransaction(ctx, trx, request.Code, balanceReduction); err != nil {
		log.Error("failed to create transaction. Err:", err)
		return errors.New("Wpooops Withdraw failed.")
	}

	return nil
}
