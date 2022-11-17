package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Ok bool `json:"ok"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Requester interface {
	MakeGetRequest(endpoint string) (ServerResponse, error)
}

type Handler struct {
	requester Requester
}

func NewHandler(requester Requester) Handler {
	return Handler{
		requester: requester,
	}
}

func (h Handler) HandleTest(c echo.Context) error {
	_, err := h.requester.MakeGetRequest("/server/test")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Ok: true})
}
