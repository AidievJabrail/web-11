package usecase

import (
	"fmt"
	"time"

	"github.com/ValeryBMSTU/web-11/pkg/vars"
	"github.com/dgrijalva/jwt-go"
)

func (uc *Usecase) GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 48).Unix(), // Токен будет действовать 72 часа
	})

	tokenString, err := token.SignedString(vars.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *Usecase) Register(username, password string) error {
	_, err := uc.provider.FoundUser(username)
	if err != nil{
		return uc.provider.CreateUser(username, password)
	}
	return fmt.Errorf("такой пользователь уже есть")
	
}

func (uc *Usecase) Login(username, password string) (string, error) {
	username, err := uc.provider.CheckUser(username, password)
	if err != nil {
		return "", err
	}

	return username, nil 
}