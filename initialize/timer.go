package initialize

import (
	"bankroll/global"
	"bankroll/service"
	"fmt"
	"time"
)

func Timer() {
	id, _ := global.Timer.AddTaskByFunc("test5", "0/59 9-16 * * 1-5", func() {
		service.MarketGetPlateBankroll(1000)
		service.MarketGetBankRoll(service.Industry, "je", 1, 100)
		service.MarketGetBankRoll(service.Conception, "je", 1, 300)
		service.MarketGetBankRoll(service.Individual, "money", 1, 5000)
	})
	global.Zlog.Warn(fmt.Sprintf("板块资金|概念资金|个股资金 定时任务Id：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test4", "30 1 1 * * ", func() {
		service.SetReportCodeToRedis()
	})
	global.Zlog.Warn(fmt.Sprintf("设置股票报告code定时任务Id：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test4", "0/5 * * * * ", func() {
		service.MarketGetStockReport()
	})
	global.Zlog.Warn(fmt.Sprintf("消费股票报告code定时任务Id：%d ", id))

	//虚拟币的行情
	id, _ = global.Timer.AddTaskByFunc("test5", "0/59 * * * *", func() {
		service.BiDealDetail()
	})
	global.Zlog.Warn(fmt.Sprintf("非小号虚拟币的交易情况：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test5", "25-29 9 * * 1-5", func() {
		for i := 0; i < 18; i++ {
			service.WenSearchBiddingData("上个交易日板块热度前4,非同花顺特色指数,非同花顺地域概念,上个交易日涨停家数>8","涨跌幅>0%且涨跌幅<30%,量比>2,委比>50,上市天数>365,a股流通市值<300亿元,非st,macd0轴上且macd>0,近一年涨幅小于50%,近10个交易日涨幅大于10%小于40%,同花顺行业,行业属于")
			time.Sleep(time.Second * 5)
		}
	})
	//service.MarketGetBankRoll(service.Conception,"je",1,300)
	//service.MarketGetBankRoll(service.Individual,"money",1,5000)

	//if global.Config.Timer.Start {
	//	for i := range global.Config.Timer.Detail {
	//		go func(detail config.Detail) {
	//			fmt.Println(detail.CompareField)
	//			global.Timer.AddTaskByFunc("test", "1s", func() {
	//				fmt.Println(time.Time{}.Date())
	//			})
	//		}(global.Config.Timer.Detail[i])
	//	}
	//}
}
