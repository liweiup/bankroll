package global

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)


type RedisStore struct {
}

func (c *RedisStore) Get(key string) ([]byte, error){
	conn := GVA_REDIS.Get()
	defer conn.Close()
	reply, err := conn.Do("get", key)
	if err != nil {
		return nil,err
	}
	return redis.Bytes(reply, err)
}
func (c *RedisStore) Set(key string, value []byte) error {
	conn := GVA_REDIS.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}
func (c *RedisStore) Exists(key string) (bool, error) {
	conn := GVA_REDIS.Get()
	defer conn.Close()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func (c *RedisStore) Delete(key string) error {
	conn := GVA_REDIS.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}
