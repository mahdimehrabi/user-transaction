package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	userController *controller.UserController
}

func NewRouteHandler(userService service.UserService) *RouteHandler {
	userController := controller.NewUserController(userService)
	return &RouteHandler{userController: userController}
}

func (rh *RouteHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/users", rh.userController.CreateUser)
	router.GET("/users/:id", rh.userController.GetUserByID)
	router.PUT("/users/:id", rh.userController.UpdateUser)
	router.DELETE("/users/:id", rh.userController.DeleteUser)
	router.GET("/users", rh.userController.GetAllUsers)
}
