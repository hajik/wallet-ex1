package http

import (
	"net/http"
	"wallet-ex1/internal/transaction"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TransactionDelivery struct {
	uc transaction.ITransactionUsecase
}

func NewTransactionDelivery(uc transaction.ITransactionUsecase) *TransactionDelivery {
	return &TransactionDelivery{uc: uc}

}

func (h *TransactionDelivery) Withdraw(c echo.Context) error {
	logger := c.Get("logger").(*zap.Logger)

	var req transaction.WithdrawRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed bind request err:", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(req); err != nil {
		logger.Error("Failed validate request err: ", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.uc.Withdraw(c.Request().Context(), req); err != nil {
		logger.Error("Failed withdraw err:", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "withdraw successful"})
}
