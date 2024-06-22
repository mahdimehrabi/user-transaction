package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/app/api/middleware"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userController controller.UserController
	authMiddleware middleware.AuthMiddleware
}

func NewUserRouter(userService service.UserService, authMiddleware middleware.AuthMiddleware) *UserRouter {
	userController := controller.NewUserController(userService)
	return &UserRouter{userController: *userController, authMiddleware: authMiddleware}
}

func (rh *UserRouter) SetupRoutes(router *gin.Engine) {
	g := router.Group("/api/users").Use(rh.authMiddleware.Handle())
	{
		g.POST("", rh.userController.CreateUser)
		g.GET("/:id", rh.userController.GetUserByID)
		g.PUT("/:id", rh.userController.UpdateUser)
		g.DELETE("/:id", rh.userController.DeleteUser)
		g.GET("", rh.userController.GetAllUsers)
	}
}
