package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
)

func (srv *Server) Register(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Не удалось прочитать данные"})
	}

	err := srv.uc.Register(input.Username, input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	token, err := srv.uc.GenerateJWT(input.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось сгенерировать токен"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Пользователь зарегистрирован!", "token": token}) 
}

func (srv *Server) Login(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Не удалось прочитать данные"})
	}

	userName, err := srv.uc.Login(input.Username, input.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Не подходящий логин или пароль"})
	}

	token, err := srv.uc.GenerateJWT(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось сгенерировать токен"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

