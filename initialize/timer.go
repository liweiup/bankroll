package initialize

import (
	"bankroll/config"
	"bankroll/global"
	"fmt"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				fmt.Println(detail.TableName)
				fmt.Println(detail.CompareField)
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {


				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}
