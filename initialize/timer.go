package initialize

import (
	"bankroll/global"
	"bankroll/service"
	"fmt"
)

func Timer() {
	id, _ := global.Timer.AddTaskByFunc("test1", "0/30 15 * * 1-5", func() {
		service.MarketGetPlateBankroll(1000)
	})
	global.Zlog.Warn(fmt.Sprintf("板块资金定时任务Id：%d ", id))
	id, _ = global.Timer.AddTaskByFunc("test2", "0/30 16 * * 1-5", func() {
		service.MarketGetBankRoll(service.Industry, "je", 1, 100)
	})
	global.Zlog.Warn(fmt.Sprintf("行业资金定时任务Id：%d ", id))
	id, _ = global.Timer.AddTaskByFunc("test3", "0/30 17 * * 1-5", func() {
		service.MarketGetBankRoll(service.Conception, "je", 1, 300)
	})
	global.Zlog.Warn(fmt.Sprintf("概念资金数据定时任务Id：%d ", id))
	id, _ = global.Timer.AddTaskByFunc("test4", "0/30 18 * * 1-5", func() {
		service.MarketGetBankRoll(service.Individual, "money", 1, 5000)
	})
	global.Zlog.Warn(fmt.Sprintf("个股资金定时任务Id：%d ", id))

	id, _ = global.Timer.AddTaskByFunc("test4", "0/59 9-15 * * 1-5", func() {
		service.MarketGetPlateBankroll(1000)
		service.MarketGetBankRoll(service.Industry, "je", 1, 100)
		service.MarketGetBankRoll(service.Conception, "je", 1, 300)
		service.MarketGetBankRoll(service.Individual, "money", 1, 5000)
	})
	//service.MarketGetBankRoll(service.Industry,"je",1,100)
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
