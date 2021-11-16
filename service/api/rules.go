package api

import "bankroll/utils"

var (
	GetPlateBankrollData  = utils.Rules{"Cdate": {utils.NotEmpty()},"CompareNum": {utils.NotEmpty()}, "PeriodNum": {utils.NotEmpty()}}
)
