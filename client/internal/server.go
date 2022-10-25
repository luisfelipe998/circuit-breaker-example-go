package internal

import (
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	handler := NewHandler(NewCircuitBreakerRequester())
	// handler := NewHandler(NewDefaultRequester())

	e := echo.New()
	e.GET("/client/ping", handler.HandlePing)
	e.GET("/client/test", handler.HandleTest)

	return e
}
