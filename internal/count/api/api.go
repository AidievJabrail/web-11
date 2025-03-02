package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Server struct {
	maxSize int

	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxSize int, uc Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		uc:      uc,
	}

	api.server = echo.New()
	api.server.GET("/count", api.JWTMiddleware(api.GetCount))
	api.server.POST("/count", api.JWTMiddleware(api.IncrementCount))
	api.server.POST("/count/set", api.JWTMiddleware(api.SetCount))

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
