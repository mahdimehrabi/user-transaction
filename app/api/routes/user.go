package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userController controller.UserController
}

func NewUserRouter(userService service.UserService) *UserRouter {
	userController := controller.NewUserController(userService)
	return &UserRouter{userController: *userController} //transient controller injection to improve performance
}

func (rh *UserRouter) SetupRoutes(router *gin.Engine) {
	g := router.Group("/api/users")
	{
		g.POST("", rh.userController.CreateUser)
		g.GET("/:id", rh.userController.GetUserByID)
		g.PUT("/:id", rh.userController.UpdateUser)
		g.DELETE("/:id", rh.userController.DeleteUser)
		g.GET("", rh.userController.GetAllUsers)
	}
}
