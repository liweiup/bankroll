package api

import (
	"bankroll/service/common"
	"bankroll/service/mapper"
)

var RedisCache = new(common.RedisCache)
var DataInfo = new(mapper.DataInfo)