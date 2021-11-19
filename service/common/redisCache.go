package common

import (
	"bankroll/service/common/response"
	"time"
)

type RedisCache struct {
}

func (rc *RedisCache) Set(key string,obj interface{},ttl time.Duration)  {

}

func (rc *RedisCache) Get(key string,wanted response.Response) error {

	//
	//x := initialize.GVA_REDIS.Set(context.Background(),key,savee,time.Hour)
	//xx := initialize.GVA_REDIS.Get(context.Background(),key)
	//json.Unmarshal([]byte(xx.Val()),ob)

	return nil
}
