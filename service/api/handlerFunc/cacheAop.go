package handlerFunc

import (
	"bankroll/config"
	"bankroll/global/redigo"
	"bankroll/service/common/response"
	"bankroll/utils"
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
func (r responseBodyWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}
func CacheAop() gin.HandlerFunc {
	fun := func(c *gin.Context) {
		//pre
		var param map[string]interface{}
		err := c.ShouldBindBodyWith(&param, binding.JSON)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		param["path"] = c.FullPath()
		mjson, _ := json.Marshal(param)
		key := fmt.Sprintf("%x", md5.Sum(mjson))
		//获取数据
		data := &response.Response{}
		str, err := redigo.Dtype.String.Get(config.CacheSet + key).String()
		if err == nil {
			str = utils.UnzipStr(str)
			err = data.UnMarshalBinary(str, data)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				c.Abort()
				return
			}
			response.OkWithDetailed(data.Data, "succ", c)
			c.Abort()
			return
		}
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
		//after 设置缓存
		err = data.UnMarshalBinary(w.body.String(), data)
		if err != nil {
			return
		}
		if data.Code == 0 {
			//redigo.Dtype.String.Set(config.CacheSet+key,utils.ZipStr(w.body.Bytes()),3600 * 3)
		}
	}
	return fun
}
