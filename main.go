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
	//fmt.Println(redigo.Dtype.String.Get("BK:CACHE:1e359f0e445df6a59d0ebc1556208b4c"))
	//fmt.Println(redigo.Dtype.String.Set("BK:CACHE:1e","122",200))
	//fmt.Println(redigo.Dtype.Set.SisMember("BK:HOLIDAY","2022-10-04").Bool())
	//var data response.Response
	//str,_ := redigo.Dtype.String.Get("BK:CACHE:4464ffcf54109c0217ac456af4377ad1").String()
	//data1 := &response.Response{}
	//fmt.Println(data1.Data)

	//fmt.Println(utils.AvgRiseRatio(0.0044,-0.0011,3))
	//fmt.Println(math.Pow(-0.25,0.333))
	//fmt.Println(utils.Sqrt(-0.239))
	//fmt.Println(math.Sqrt(-0.239))

}

