package jwt

import (
	"bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"bbdk/utils/encrypt"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	env            *godotenv.Env
	logger         logger.Logger
	userRepository userRepo.Repository
}

func NewAuthService(
	env *godotenv.Env,
	logger logger.Logger,
	userRepository userRepo.Repository,
) *AuthService {
	return &AuthService{
		env:            env,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (s AuthService) CreateAccessToken(user *entity.User, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{
		"authorized": true,
		"userID":     user.ID,
		"email":      user.Email,
		"exp":        exp,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthService) CreateRefreshToken(user entity.User, exp int64, secret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = exp
	atClaims["userID"] = user.ID
	atClaims["email"] = user.Email
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s AuthService) CreateTokens(user *entity.User, remember bool) (map[string]string, error) {
	accessSecret := "access" + os.Getenv("Secret")
	expAccessToken := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := s.CreateAccessToken(user, expAccessToken, accessSecret)

	refreshSecret := "refresh" + os.Getenv("Secret")
	var expRefreshToken int64
	if remember {
		expRefreshToken = time.Now().Add(time.Hour * 24 * 15).Unix()
	} else {
		expRefreshToken = time.Now().Add(time.Hour * 24).Unix()
	}
	refreshToken, err := s.CreateRefreshToken(*user, expRefreshToken, refreshSecret)

	return map[string]string{
		"refreshToken":    refreshToken,
		"accessToken":     accessToken,
		"expRefreshToken": strconv.Itoa(int(expRefreshToken)),
		"expAccessToken":  strconv.Itoa(int(expAccessToken)),
	}, err
}

func (s AuthService) Login(email, enteredPassword string, remember bool) (user *entity.User, tokensData map[string]string, err error) {
	user, err = s.userRepository.FindByField("email", email)
	if errors.Is(err, userRepo.ErrNotFound) {
		return
	}
	if err != nil {
		s.logger.Errorf("Error to find user:%s", err.Error())
		return
	}

	encryptedPassword := encrypt.HashSHA256(enteredPassword)
	if user.Password == encryptedPassword {
		tokensData, err = s.CreateTokens(user, remember)
		if err != nil {
			s.logger.Errorf("Failed generate jwt tokens:%s", err.Error())
			return
		}
		return
	} else {
		err = userRepo.ErrNotFound
		return
	}
}

func (s AuthService) RenewToken(refreshToken string) (accessToken string, expAccessToken int64, err error) {
	var valid bool
	var atClaims jwt.MapClaims
	refreshSecret := "refresh" + s.env.Secret
	valid, atClaims, err = DecodeToken(refreshToken, refreshSecret)
	var vErr *jwt.ValidationError
	if errors.As(err, &vErr) {
		err = userRepo.ErrNotFound
		return
	}
	if err != nil {
		return
	}

	uid, ok := atClaims["userID"].(float64)
	if !ok {
		return
	}

	user, err := s.userRepository.FindByField("id", uid)
	//don't allow deleted user renew access token
	if errors.Is(err, userRepo.ErrNotFound) {
		return
	}
	if err != nil {
		s.logger.Errorf("error in finding user:%s", err)
		return
	}

	if valid {
		expAccessToken = time.Now().Add(time.Minute * 30).Unix()
		accessToken, err = s.CreateAccessToken(user, expAccessToken, s.env.Secret)
		if err != nil {
			s.logger.Errorf("error in creating access token:%s", err)
			return
		}
		return
	}
	err = userRepo.ErrNotFound
	return
}
