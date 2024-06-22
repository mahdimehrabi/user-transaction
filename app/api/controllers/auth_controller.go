package controller

import (
	"bbdk/app/api/dto"
	"bbdk/app/api/jwt"
	"bbdk/app/api/response"
	userRepo "bbdk/domain/repository/user"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController struct {
	logger      logger.Logger
	env         *godotenv.Env
	authService *jwt.AuthService
}

func NewAuthController(
	env *godotenv.Env,
	authService *jwt.AuthService,
) *AuthController {
	return &AuthController{
		env:         env,
		authService: authService,
	}
}

func (ac AuthController) Login(c *gin.Context) {
	// Data Parse
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}
	user, tokensData, err := ac.authService.Login(loginRequest.Email, loginRequest.Password,
		loginRequest.Remember)
	if errors.Is(err, userRepo.ErrNotFound) {
		response.Response(c, gin.H{}, http.StatusNotFound, "No users found with entered credentials")
		return
	}
	if err != nil {
		response.InternalServerError(c)
		return
	}

	var loginResult dto.LoginResponse
	loginResult.AccessToken = tokensData["accessToken"]
	loginResult.RefreshToken = tokensData["refreshToken"]
	loginResult.ExpRefreshToken = tokensData["expRefreshToken"]
	loginResult.ExpAccessToken = tokensData["expAccessToken"]
	var userResponse dto.UserResponse
	userResponse.FromEntity(user)
	loginResult.User = userResponse
	response.Response(c, loginResult, http.StatusOK, "")
	return

}

func (ac AuthController) AccessTokenVerify(c *gin.Context) {
	at := dto.AccessTokenReq{}
	if err := c.ShouldBindJSON(&at); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	accessToken := at.AccessToken
	accessSecret := "access" + ac.env.Secret
	valid, _, err := jwt.DecodeToken(accessToken, accessSecret)
	if err != nil {
		response.Response(c, gin.H{}, http.StatusUnauthorized, "access token is not valid")
		return
	}

	if valid {
		response.Response(c, gin.H{}, http.StatusOK, "access token is valid")
		return
	} else {
		response.Response(c, gin.H{}, http.StatusUnauthorized, "access token is not valid")
		return
	}
}

func (ac AuthController) RenewToken(c *gin.Context) {
	rtr := dto.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&rtr); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, exp, err := ac.authService.RenewToken(rtr.RefreshToken)
	if errors.Is(err, userRepo.ErrNotFound) {
		response.Response(c, gin.H{}, http.StatusBadRequest, "access token is not valid")
		return
	}
	if err != nil {
		response.InternalServerError(c)
		return
	}
	response.Response(c,
		dto.AccessTokenRes{AccessToken: accessToken,
			ExpAccessToken: strconv.Itoa(int(exp))},
		http.StatusOK, "")
}
