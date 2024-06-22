package routes

import (
	"bbdk/app/api/jwt"
	"bbdk/app/api/middleware"
	gormUserRepo "bbdk/domain/repository/user/gorm"
	"bbdk/domain/service"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Router interface {
	SetupRoutes(engine *gin.Engine)
}

func CreateRouters(env *godotenv.Env, logger logger.Logger) []Router {
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to database error:%s", err.Error())
	}
	userRepo := gormUserRepo.NewUserRepository(db)
	userService := service.NewUserService(userRepo, logger)
	authService := jwt.NewAuthService(env, logger, userRepo)
	authMiddleware := middleware.NewAuthMiddleware(logger, env)

	return []Router{NewUserRouter(userService), NewAuthRouter(env, authService, authMiddleware)}
}

func HandleRouters(e *gin.Engine, routers []Router) {
	for _, router := range routers {
		router.SetupRoutes(e)
	}
}
