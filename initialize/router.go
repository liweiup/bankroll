package initialize

import (
	_ "bankroll/docs"
	"bankroll/global"
	"bankroll/service/api"
	"bankroll/service/api/handlerFunc"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)



var apiRouterApi = new(api.BankrollApi)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// 跨域
	Router.Use(Cors()) // 如需跨域可以打开
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Group("api").POST("getPlateBankrollData",handlerFunc.CacheAop(),apiRouterApi.GetPlateBankrollData)
	return Router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
