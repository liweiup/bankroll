package initialize

func Timer() {
	//id, _ := global.Timer.AddTaskByFunc("test5", "0/59 9-16 * * 1-5", func() {
	//	service.MarketGetPlateBankroll(1000)
	//	service.MarketGetBankRoll(service.Industry, "je", 1, 100)
	//	service.MarketGetBankRoll(service.Conception, "je", 1, 300)
	//	service.MarketGetBankRoll(service.Individual, "money", 1, 5000)
	//})
	//global.Zlog.Warn(fmt.Sprintf("板块资金|概念资金|个股资金 定时任务Id：%d ", id))
	//
	//id, _ = global.Timer.AddTaskByFunc("test4", "30 1 1 * * ", func() {
	//	service.SetReportCodeToRedis()
	//})
	//global.Zlog.Warn(fmt.Sprintf("设置股票报告code定时任务Id：%d ", id))
	//
	//id, _ = global.Timer.AddTaskByFunc("test4", "0/5 * * * * ", func() {
	//	service.MarketGetStockReport()
	//})
	//global.Zlog.Warn(fmt.Sprintf("消费股票报告code定时任务Id：%d ", id))


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
