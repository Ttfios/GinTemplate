package server

import (
	Config "ttemplate/conf"
	"ttemplate/controllers"
	"ttemplate/middleware"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(Config.Configuration.Server.GinMode)
	r := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", Config.Configuration.Redis.Address, Config.Configuration.Redis.Password, []byte("secret"))

	r.Use(sessions.Sessions("tfios-templete", store))

	r.Use(middleware.CORSMiddleware())

	r.GET("/", controllers.Ping)

	return r
}
