package api

import (
	"bbdk/infrastructure/godotenv"
	"bbdk/infrastructure/log/zerolog"
	"github.com/gin-gonic/gin"
)

func Boot() {
	r := gin.Default()
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()

	//I must define router struct but for lack of time I call handler(controller) directly
	router(r, addressUser)

	r.Run()
}

func router(r *gin.Engine, addressUser interface{}) {
	r.POST("/api/users", addressUser.CreateUser)
	r.GET("/api/users/:id", addressUser.DetailUser)
}
