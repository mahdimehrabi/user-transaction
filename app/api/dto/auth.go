package dto

import "bbdk/domain/entity"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remember bool   `json:"remember"`
}

func (l LoginRequest) ToUser() *entity.User {
	return &entity.User{
		Email:    l.Email,
		Password: l.Password,
	}
}

type LoginResponse struct {
	AccessToken     string       `json:"accessToken"`
	RefreshToken    string       `json:"refreshToken"`
	ExpRefreshToken string       `json:"expRefreshToken"`
	ExpAccessToken  string       `json:"expAccessToken"`
	User            UserResponse `json:"user"`
}

type AccessTokenReq struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

type AccessTokenRes struct {
	AccessToken    string `json:"accessToken" binding:"required"`
	ExpAccessToken string `json:"expAccessToken" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type TokenRequest struct {
	Token string `json:"token" binding:"len=40,required"`
}

type TokenRequestNoLimit struct {
	Token string `json:"token" binding:"required"`
}
