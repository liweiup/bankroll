package main

import (
	"bankroll/global"
	redigo_pack "bankroll/global/redigo"
	"bankroll/initialize"
	"fmt"
)

func main() {
	initialize.Viper()
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	global.GVA_LOG = initialize.Zap()
	initialize.Redis()
	initialize.Timer()
	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	//initialize.RunWindowsServer()
	//var p *response.Response
	//var r = new(global.RedisStore)
	//api.RedisCache.Get("BK:CACHE:1e359f0e445df6a59d0ebc1556208b4c",p)
	//r.Get("BK:CACHE:1e359f0e445df6a59d0ebc1556208b4c")
	fmt.Println(redigo_pack.RedigoConn.String.Get("BK:CACHE:1e359f0e445df6a59d0ebc1556208b4c"))
	//var redigo =
}

