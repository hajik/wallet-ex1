package http

import (
	"net/http"
	"strconv"
	"wallet-ex1/internal/wallet"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type WalletDelivery struct {
	uc wallet.IWalletUsecase
}

func NewWalletDelivery(uc wallet.IWalletUsecase) *WalletDelivery {
	return &WalletDelivery{uc: uc}

}

func (h *WalletDelivery) GetBalance(c echo.Context) error {
	logger := c.Get("logger").(*zap.Logger)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := h.uc.GetWallet(c.Request().Context(), id)
	if err != nil {
		logger.Error("Failed get balance. Error is", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": result})
}
