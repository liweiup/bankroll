package initialize

import (
	"bankroll/global"
	"bankroll/service"
	"fmt"
)

func Timer() {
	id, _ := global.GVA_Timer.AddTaskByFunc("test", "0/20 16 * * 1-5", func() {
	//id, _ := global.GVA_Timer.AddTaskByFunc("test", "0/20 16 * * ? ", func() {
		service.MarketGetBankRoll(service.Industry,"je",1,100)
	})
	global.GVA_LOG.Warn(fmt.Sprintf("Industry添加定时任务Id：%d ",id))
	id, _ = global.GVA_Timer.AddTaskByFunc("test", "0/20 17 * * 1-5", func() {
		service.MarketGetBankRoll(service.Conception,"je",1,100)
	})
	global.GVA_LOG.Warn(fmt.Sprintf("Conception添加定时任务Id：%d ",id))
	id, _ = global.GVA_Timer.AddTaskByFunc("test", "0/20 18 * * 1-5", func() {
		service.MarketGetBankRoll(service.Individual,"money",1,5000)
	})
	global.GVA_LOG.Warn(fmt.Sprintf("Individual添加定时任务Id：%d ",id))
	//service.MarketGetBankRoll(service.Industry,"je",1,100)
	//service.MarketGetBankRoll(service.Conception,"je",1,100)
	//service.MarketGetBankRoll(service.Individual,"money",1,5000)
	//if global.GVA_CONFIG.Timer.Start {
	//	for i := range global.GVA_CONFIG.Timer.Detail {
	//		go func(detail config.Detail) {
	//			fmt.Println(detail.CompareField)
	//			global.GVA_Timer.AddTaskByFunc("test", "1s", func() {
	//				fmt.Println(time.Time{}.Date())
	//			})
	//		}(global.GVA_CONFIG.Timer.Detail[i])
	//	}
	//}
}
