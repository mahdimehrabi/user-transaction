package middleware

import (
	jwt2 "bbdk/app/api/jwt"
	"bbdk/app/api/response"
	"bbdk/domain/entity"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

// AuthMiddleware -> struct for transaction
type AuthMiddleware struct {
	logger logger.Logger
	env    *godotenv.Env
}

// NewAuthMiddleware -> new instance of transaction
func NewAuthMiddleware(
	logger logger.Logger,
	env *godotenv.Env,
) *AuthMiddleware {
	return &AuthMiddleware{
		logger: logger,
		env:    env,
	}
}

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {

	return func(c *gin.Context) {
		ah := authHeader{}
		if err := c.ShouldBindHeader(&ah); err == nil {
			strs := strings.Split(ah.Authorization, " ")
			if len(strs) != 2 {
				response.Response(c, nil, http.StatusUnauthorized, "your access token is not correct")
				c.Abort()
				return
			}
			bearer := strs[0]
			if bearer != "Bearer" {
				response.Response(c, nil, http.StatusUnauthorized, "your token doesn't start with 'Bearer '")
				c.Abort()
				return
			}
			accessToken := strs[1]
			valid, claims, _ := jwt2.DecodeToken(accessToken, "access"+m.env.Secret)
			user, ok := m.claimsToUser(claims)
			if !ok {
				response.Response(c, nil, http.StatusUnauthorized, "You must login to access this page 😥")
				c.Abort()
				return
			}
			if valid && err == nil {
				c.Set("user", user)
				c.Next()
				return
			}
		}
		response.Response(c, nil, http.StatusUnauthorized, "You must login to access this page 😥")
		c.Abort()
	}
}

// claimsToUser convert jwt claims to user model
func (m AuthMiddleware) claimsToUser(claims jwt.MapClaims) (user *entity.User, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	user = new(entity.User)
	id := claims["userID"].(float64)
	user.ID = uint(id)
	user.Email = claims["email"].(string)
	ok = true
	return
}

// AuthenticatedUser return authenticated user from gin context that filled by jwt claims
// if no user stored in gin context it returns empty user
func AuthenticatedUser(c *gin.Context) *entity.User {
	user := func() *entity.User {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		return c.MustGet("user").(*entity.User)
	}()
	if user == nil {
		user = &entity.User{}
	}
	return user
}
