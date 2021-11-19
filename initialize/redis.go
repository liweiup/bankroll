package initialize

import (
	"bankroll/global"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"time"
)

func Redis() {
	redisCfg := global.Config.Redis
	pool := &redis.Pool{
		MaxActive: 512,
		MaxIdle: 10,
		Wait: false,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			c,err := redis.Dial("tcp",redisCfg.Addr,redis.DialPassword(redisCfg.Password),redis.DialDatabase(2))
			if err != nil {
				global.Zlog.Error("redis connect failed, err:", zap.Any("err", err))
				return c,err
			}
			return c,err
		},
	}
	global.Redigo = pool
}