package model

type AuthForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
