package internal

import (
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	handler := NewHandler(NewDefaultRequester())
	// handler := NewHandler(NewCircuitBreakerRequester())
	// handler := NewHandler(NewCustomCircuitBreakerRequester())

	e := echo.New()
	e.GET("/client/test", handler.HandleTest)

	return e
}
