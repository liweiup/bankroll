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
	//for i := 0; i < 5; i++ {
	//service.WenSearchBiddingData("上个交易日板块热度前10；非同花顺特色指数；非同花顺地域概念","涨跌幅大于0%且涨跌幅小于30%；量比大于2；委比大于50%；上市天数大于1年；流通市值小于300亿；非st的股票；macd零轴上；近1年涨跌幅小于50%；近10个交易日涨幅大于10%且小于40%；同花顺行业；不包含创业板股票；")
	//service.WenSearchBiddingData("上个交易日板块热度前4,非同花顺特色指数,非同花顺地域概念,上个交易日涨停家数>8","涨跌幅>0%且涨跌幅<30%,量比>2,上市天数>365,a股流通市值<300亿元,非st,macd0轴上且macd>0,近一年涨幅小于50%,近10个交易日涨幅大于10%小于40%,同花顺行业,行业属于")
	//}
	//initialize.RunWindowsServer()
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
