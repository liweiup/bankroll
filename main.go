package main

import (
	"bankroll/global"
	"bankroll/initialize"
	"bankroll/service"
)

func main() {
	initialize.Viper()
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	service.MarketGetBankRoll("gnzjl",1,5)
	//println(global.GVA_DB)
	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
}


