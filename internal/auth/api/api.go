package api

import (
	"fmt"

	"github.com/ValeryBMSTU/web-11/internal/auth/usecase"
	"github.com/labstack/echo/v4"
)


type Server struct {
	Address string
	Router  *echo.Echo
	uc      *usecase.Usecase
}

func NewServer(ip string, port int, uc *usecase.Usecase) *Server {
	e := echo.New()
	srv := &Server{
		Address: fmt.Sprintf("%s:%d", ip, port),
		Router:  e,
		uc:      uc,
	}

	srv.Router.POST("/auth/register", srv.Register)
	srv.Router.POST("/auth/login", srv.Login)

	return srv
}

