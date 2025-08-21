package sdk

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator adalah struct yang mengimplementasikan echo.Validator.
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate adalah metode yang diperlukan oleh antarmuka echo.Validator.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Mengembalikan error yang terstruktur agar respons lebih jelas.
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
