package main

import (
	"bankroll/global"
	"bankroll/initialize"
)

func main() {
	initialize.Viper()
	global.Gdb = initialize.Gorm() // gorm连接数据库
	global.Zlog = initialize.Zap()
	initialize.Redis()
	initialize.Timer()
	if global.Gdb != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.Gdb.DB()
		defer db.Close()
	}
	initialize.RunWindowsServer()
	//service.SetReportCodeToRedis()
	//service.MarketGetStockReport()
	//reportList, _ := api.StockReport.GetIndividualReport("002739","")
	//avgRiseRatio, _ := utils.AvgRiseRatio(reportList[len(reportList)-1].RetainProfit, reportList[0].RetainProfit, float64(len(reportList)-1))
	//fmt.Println(avgRiseRatio)

	//fmt.Println(utils.GetPeriodByOneday("2022-01-05", 1, 4))

	//fmt.Println(redigo.Dtype.Set.SisMember("BK:HOLIDAY","2022-10-04").Bool())
	//var data response.Response
	//str,_ := redigo.Dtype.String.Get("BK:CACHE:4464ffcf54109c0217ac456af4377ad1").String()
	//data1 := &response.Response{}
	//fmt.Println(data1.Data)
	//fmt.Println(utils.AvgRiseRatio(0.0044,-0.0011,3))
	//fmt.Println(math.Pow(-0.25,0.333))
	//fmt.Println(utils.Sqrt(-0.239))
	//fmt.Println(math.Sqrt(-0.239))
	//fmt.Println("binarySearchRecursive target index: ", binarySearchRecursive(s,9, 0, len(s)))
}
