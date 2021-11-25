package initialize

import (
	_ "bankroll/docs"
	"bankroll/global"
	"bankroll/service/api"
	"bankroll/service/api/handlerFunc"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)



var apiRouterApi = new(api.BankrollApi)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	//跨域
	Router.Use(Cors(),GzipHandler().Gin) // 如需跨域可以打开
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Group("api").POST("getPlateBankrollData",handlerFunc.CacheAop(),apiRouterApi.GetPlateBankrollData)
	Router.Group("api").POST("getStockBankrollData",handlerFunc.CacheAop(),apiRouterApi.GetStockBankrollData)
	Router.Group("api").POST("getPlateGroup",handlerFunc.CacheAop(),apiRouterApi.GetPlateGroup)
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

//gzip handler
func GzipHandler() *gzip.Handler  {
	handler := gzip.NewHandler(gzip.Config{
		// gzip压缩等级
		CompressionLevel: 6,
		// 触发gzip的最小body体积，单位：byte
		MinContentLength: 1024,
		// 请求过滤器基于请求来判断是否对这条请求的返回启用gzip，
		// 过滤器按其定义顺序执行，下同。
		RequestFilter: []gzip.RequestFilter{
			gzip.NewCommonRequestFilter(),
			gzip.DefaultExtensionFilter(),
		},
		// 返回header过滤器基于返回的header判断是否对这条请求的返回启用gzip
		ResponseHeaderFilter: []gzip.ResponseHeaderFilter{
			gzip.NewSkipCompressedFilter(),
			gzip.DefaultContentTypeFilter(),
		},
	})
	return handler
}