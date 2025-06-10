package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pusrenk/auth-service/internal/grpc" // adjust if the import path is different
)

func main() {
	// Initialize gRPC clients (e.g., customer service)
	grpc.Init()

	// Set up Echo server
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
