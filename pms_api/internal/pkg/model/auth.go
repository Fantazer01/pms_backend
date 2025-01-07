package model

import "github.com/golang-jwt/jwt/v5"

type AuthForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AppClaims struct {
	Username       string `json:"name"`
	Admin          bool   `json:"admin"`
	IsRefreshToken bool   `json:"is_refresh_token"`
	jwt.RegisteredClaims
}
