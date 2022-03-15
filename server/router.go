package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goworkeryyt/go-core/global"
	."granada1/server/controllers/auth"
	"granada1/server/controllers/table"
	"granada1/server/middlewares"
)

func InitApiRouter(r *gin.Engine) *gin.Engine{
	Cors(r)
	//Gin 默认信任所有代理，这是不安全的。同时，如果您不使用任何代理，您可以使用禁用此功能Engine.SetTrustedProxies(nil)
	err := r.SetTrustedProxies(nil)
	if err!=nil {
		global.LOG.Error(err.Error())
	}

	//TODO:接口
	auth:=r.Group("/auth")
	{
		auth.GET("/code/:account",SendCode)
		auth.POST("/register",Register)
		auth.POST("login",Login)

	}
	app:=r.Group("/app",middlewares.AuthMiddleWare())
	{
		app.POST("/user/info",table.AddInfo)
		app.POST("/user/address",table.AddAddress)
		app.POST("/user/community",table.CreateCommunity)
		app.POST("/user/relation",table.CreateRelation)
		app.GET("/user/ids",table.SelectAllUserID)
	}

	return r
}

func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", "range")
	r.Use(cors.New(config))
}

