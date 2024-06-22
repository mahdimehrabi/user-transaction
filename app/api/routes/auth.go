package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/app/api/jwt"
	"bbdk/app/api/middleware"
	"bbdk/infrastructure/godotenv"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController *controller.AuthController
}

func NewAuthRouter(env *godotenv.Env,
	authService *jwt.AuthService,
	authMiddleware *middleware.AuthMiddleware) *AuthRouter {
	authController := controller.NewAuthController(env, authService, authMiddleware)
	return &AuthRouter{authController: authController} //transient controller injection to improve performance
}

func (rh *AuthRouter) SetupRoutes(router *gin.Engine) {
	g := router.Group("/api/auth")
	{
		g.POST("/login", rh.authController.Login)
		g.POST("/access-token-verify", rh.authController.AccessTokenVerify)
		g.POST("/renew-access-token", rh.authController.RenewToken)
	}
}
