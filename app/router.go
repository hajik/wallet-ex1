package app

import (
	"net/http"

	trxDelivery "wallet-ex1/internal/transaction/delivery"
	trxtUsecase "wallet-ex1/internal/transaction/usecase"
	walletDelivery "wallet-ex1/internal/wallet/delivery"
	walletUsecase "wallet-ex1/internal/wallet/usecase"

	"github.com/labstack/echo/v4"
)

func (s *server) initRoutes() {

	s.server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	r := repo{}
	r.InitRepository(s.db)

	// authUc wallet UseCase
	walletUC := walletUsecase.NewWalletUsecase(r.walletRepository)
	trxUC := trxtUsecase.NewTransactionUsecase(r.trxRepository)

	// Delivery
	walletHd := walletDelivery.NewWalletDelivery(walletUC)
	trxHd := trxDelivery.NewTransactionDelivery(trxUC)

	// Initialize route group
	apiRoute := s.server.Group("/api/v1")

	// Group wallet
	walletRoute := apiRoute.Group("/wallets")
	walletRoute.GET("/balance/:id", walletHd.GetBalance)

	// Group transaction
	trxRoute := apiRoute.Group("/transactions")
	trxRoute.POST("/withdraw", trxHd.Withdraw)

}
