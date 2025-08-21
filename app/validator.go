package app

import (
	"wallet-ex1/app/sdk"

	"github.com/go-playground/validator/v10"
)

func (s *server) initValidator() {

	// Register our custom validator to echo framework
	s.server.Validator = &sdk.CustomValidator{Validator: validator.New()}
}
