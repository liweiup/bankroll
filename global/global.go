package global

import (
	"bankroll/config"
	"bankroll/utils/timer"
	"github.com/garyburd/redigo/redis"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Gdb     *gorm.DB
	Redigo  *redis.Pool
	Timer  timer.Timer = timer.NewTimerTask()
	Config config.Server
	Viper     *viper.Viper
	Zlog    *zap.Logger
	Bcache local_cache.Cache
)
