package routes

import (
	"bbdk/app/api/jwt"
	"bbdk/app/api/middleware"
	gormTransRepo "bbdk/domain/repository/transaction/gorm"
	gormUserRepo "bbdk/domain/repository/user/gorm"
	"bbdk/domain/service"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"bbdk/utils/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Router interface {
	SetupRoutes(engine *gin.Engine)
}

func CreateRouters(env *godotenv.Env, logger logger.Logger) []Router {
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	fk := validator.NewFkValidator(db)
	if err := fk.Setup(); err != nil {
		logger.Fatalf("failed to setup fk validator error:%s", err.Error())
	}
	if err != nil {
		logger.Fatalf("failed to connect to database error:%s", err.Error())
	}
	userRepo := gormUserRepo.NewUserRepository(db)
	transRepo := gormTransRepo.NewTransactionRepository(db)
	userService := service.NewUserService(userRepo, logger)
	transactionService := service.NewTransactionService(transRepo, logger)
	authService := jwt.NewAuthService(env, logger, userRepo)
	authMiddleware := middleware.NewAuthMiddleware(logger, env)

	return []Router{NewAuthRouter(env, authService),
		NewUserRouter(userService, *authMiddleware), NewTransactionRouter(transactionService, *authMiddleware)}
}

func HandleRouters(e *gin.Engine, routers []Router) {
	for _, router := range routers {
		router.SetupRoutes(e)
	}
}
