package main

import (
	"bankroll/global"
	"bankroll/initialize"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

	//双色球
	//篮球
	blueBall := []int{1,2,3,4,5,6,7,8,9,10,11,12,14,14,15,16}
	//红球
	redBall := []int{1,2,3,4,5,6,7,8,9,10,11,12,14,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33}
	//取6个红球
	flagBall := []string{}
	rand.Seed(time.Now().Unix())
	for{
		br := rand.Intn(32)  //随机数
		forflag := false
		for i2, _ := range flagBall {
			if i2 == redBall[br] {
				forflag = true
			}
		}
		if forflag {
			continue
		}
		fmt.Println(br)
		if redBall[br] != 0 {
			flagBall = append(flagBall, strconv.Itoa(redBall[br]))
		}
		if len(flagBall) == 6 {
			break
		}
	}
	rr := rand.Intn(15)
	fmt.Println(strings.Join(flagBall," ") + " " + strconv.Itoa(blueBall[rr]))
}
