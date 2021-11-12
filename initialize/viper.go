package initialize

import (
	"bankroll/config"
	"bankroll/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"time"
)

func Viper(path ...string) *viper.Viper {
	var conf string
	if len(path) == 0 {
		flag.StringVar(&conf,"c","","choose config file.")
		flag.Parse()
		if conf == "" {
			conf = config.ConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", conf)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", conf)
		}
	}
	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(3600)),
	)
	// 将文件内容解析后封装到global对象中
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}