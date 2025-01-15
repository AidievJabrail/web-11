package api

import (
	"errors"
	"github.com/ValeryBMSTU/web-11/pkg/vars"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func (s *Server) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is required"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return vars.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Зарегиструйтесь через /auth"})
		}

		return next(c)
	}
}

func (srv *Server) GetHello(e echo.Context) error {
	ip := e.Request().RemoteAddr
	msg, err := srv.uc.FetchHelloMessage(ip)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}


func (srv *Server) PostUser(e echo.Context) error {
	ip := e.Request().RemoteAddr
	input := struct {
		Msg *string `json:"msg"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "name is empty")
	}

	if len([]rune(*input.Msg)) > srv.maxSize {
		return e.String(http.StatusBadRequest, "name too large")
	}

	err = srv.uc.SetHelloMessage(*input.Msg, ip)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusCreated, "OK")
}
