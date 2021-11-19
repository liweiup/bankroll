package handlerFunc

import (
	"bankroll/config"
	"bankroll/global/redigo"
	"bankroll/service/common/response"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		_,err := redigo.Dtype.String.Get(config.CacheSet+key).String()
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
		redigo.Dtype.String.Set(config.CacheSet+key,w.body.String())
	}
	return fun
}
