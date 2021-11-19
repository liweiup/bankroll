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
	//initialize.RunWindowsServer()
	//fmt.Println(redigo.Dtype.String.Get("BK:CACHE:1e359f0e445df6a59d0ebc1556208b4c"))
	//fmt.Println(redigo.Dtype.String.Set("BK:CACHE:1e","122",200))
	//fmt.Println(redigo.Dtype.Set.SisMember("BK:HOLIDAY","2022-10-04").Bool())
}

