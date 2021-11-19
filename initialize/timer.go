package initialize

import (
	"bankroll/global"
	"bankroll/service"
	"fmt"
)

func Timer() {
	id, _ := global.Timer.AddTaskByFunc("test", "0/20 16 * * 1-5", func() {
		service.MarketGetBankRoll(service.Industry,"je",1,100)
	})
	global.Zlog.Warn(fmt.Sprintf("Industry添加定时任务Id：%d ",id))
	id, _ = global.Timer.AddTaskByFunc("test", "0/20 17 * * 1-5", func() {
		service.MarketGetBankRoll(service.Conception,"je",1,300)
	})
	global.Zlog.Warn(fmt.Sprintf("Conception添加定时任务Id：%d ",id))
	id, _ = global.Timer.AddTaskByFunc("test", "0/20 18 * * 1-5", func() {
		service.MarketGetBankRoll(service.Individual,"money",1,5000)
	})
	global.Zlog.Warn(fmt.Sprintf("Individual添加定时任务Id：%d ",id))
	//service.MarketGetBankRoll(service.Industry,"je",1,100)
	//service.MarketGetBankRoll(service.Conception,"je",1,100)
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
