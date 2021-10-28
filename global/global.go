package global

import (
	"bankroll/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	BlackCache local_cache.Cache
)
