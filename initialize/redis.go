package initialize

import (
	"bankroll/global"
	"fmt"

	//"context"
	//"github.com/go-redis/redis/v8"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"time"
)
//
//func Redis() {
//	redisCfg := global.GVA_CONFIG.Redis
//	client := redis.NewClient(&redis.Options{
//		Addr:     redisCfg.Addr,
//		Password: redisCfg.Password, // no password set
//		DB:       redisCfg.DB,       // use default DB
//	})
//	pong, err := client.Ping(context.Background()).Result()
//	if err != nil {
//		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
//	} else {
//		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
//		global.GVA_REDIS = client
//	}
//}


func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	pool := &redis.Pool{
		MaxActive: 512,
		MaxIdle: 10,
		Wait: false,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			c,err := redis.Dial("tcp",redisCfg.Addr,redis.DialPassword(redisCfg.Password),redis.DialDatabase(2))
			if err != nil {
				global.GVA_LOG.Error("redis connect failed, err:", zap.Any("err", err))
				return c,err
			}
			return c,err
		},
	}
	fmt.Println(pool)
	fmt.Println(111)
	global.GVA_REDIS = pool
	fmt.Println(global.GVA_REDIS)
}