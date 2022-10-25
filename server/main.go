package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Status struct {
	Up bool `json:"up"`
}

type Response struct {
	Ok bool `json:"ok"`
}

type Error struct {
	Error string `json:"error"`
}

var up = true

func main() {
	e := echo.New()
	e.GET("/server/test", func(c echo.Context) error {
		if !up {
			return c.JSON(http.StatusInternalServerError, Response{Ok: false})
		}
		return c.JSON(http.StatusOK, Response{Ok: true})
	})

	e.PUT("/server/status", func(c echo.Context) error {
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{Error: err.Error()})
		}
		var status Status
		err = json.Unmarshal(bodyBytes, &status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{Error: err.Error()})
		}

		up = status.Up
		return c.JSON(http.StatusOK, status)
	})

	if os.Getenv("PORT") == "" {
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
	}
}
