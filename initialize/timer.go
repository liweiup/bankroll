package initialize

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/service"
	"bankroll/utils"
	"fmt"
	"time"
)

func Timer() {
	id, _ := global.Timer.AddTaskByFunc("test1", "0/59 9-16 * * 1-5", func() {
		service.MarketGetPlateBankroll(1000)
		service.MarketGetBankRoll(service.Industry, "je", 1, 100)
		service.MarketGetBankRoll(service.Conception, "je", 1, 300)
		service.MarketGetBankRoll(service.Individual, "money", 1, 5000)
	})
	global.Zlog.Warn(fmt.Sprintf("板块资金|概念资金|个股资金 定时任务Id：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test2", "30 1 1 * * ", func() {
		service.SetReportCodeToRedis()
	})
	global.Zlog.Warn(fmt.Sprintf("设置股票报告code定时任务Id：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test3", "0/5 * * * * ", func() {
		service.MarketGetStockReport()
	})
	global.Zlog.Warn(fmt.Sprintf("消费股票报告code定时任务Id：%d ", id))

	//虚拟币的行情
	id, _ = global.Timer.AddTaskByFunc("test4", "0/59 * * * *", func() {
		service.BiDealDetail()
	})
	global.Zlog.Warn(fmt.Sprintf("非小号虚拟币的交易情况：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test5", "25-29 9 * * 1-5", func() {
		for i := 0; i < 18; i++ {
			service.WenSearchBiddingData("上个交易日板块热度前4；非同花顺特色指数；非同花顺地域概念；上个交易日涨停家数大于8","涨跌幅大于0%且涨跌幅小于30%；量比大于2；委比大于50%；上市天数大于1年；流通市值小于300亿；非st的股票；macd零轴上；近1年涨跌幅小于50%；近10个交易日涨幅大于10%且小于40%；同花顺行业；不包含创业板股票；")
			time.Sleep(time.Second * 5)
		}
	})
	id, _ = global.Timer.AddTaskByFunc("test6", "20 15 * * 1-5", func() {
		service.WenSearchLongHuData("龙虎榜；不含科创板；不含创业板；")
	})
	//service.MarketGetBankRoll(service.Conception,"je",1,300)
	//service.MarketGetBankRoll(service.Individual,"money",1,5000)
	if global.Config.Timer.Start {
		for i := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("test", "20 1 1 * *", func() {
					utils.ClearTable(global.Gdb,detail.TableName,detail.CompareField,detail.Interval)
				})
			}(global.Config.Timer.Detail[i])
		}
	}
}
