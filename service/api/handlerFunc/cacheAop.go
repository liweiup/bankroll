package handlerFunc

import (
	"bankroll/config"
	"bankroll/service/api"
	"bankroll/service/common/response"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error)  {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}
func CacheAop() gin.HandlerFunc {
	fun := func(c *gin.Context) {
		//pre
		var param map[string]interface{}
		_ = c.ShouldBindBodyWith(&param,binding.JSON)
		param["path"] = c.FullPath()
		mjson,_ := json.Marshal(param)
		key := fmt.Sprintf("%x", md5.Sum(mjson))
		//获取数据
		var data response.Response
		err := api.RedisCache.Get(config.CacheSet+key,data)
		//如果存在数据
		if err == nil {
			response.OkWithDetailed(data,"succ",c)
			c.Abort()
			return
		}
		//获取key
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
		//after 设置缓存
		api.RedisCache.Set(config.CacheSet+key,w.body.String(),time.Hour * 2)
	}
	return fun
}
