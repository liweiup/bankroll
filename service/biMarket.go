package service

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/service/model"
	"bankroll/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (

)

/**
非小号行情
 */
func BiDealDetail() {
	cDate := time.Now().Format(config.DayOut)
	url := "https://dncapi.fxhapp.com/api/coin/web-coinrank?page=1&type=-1&pagesize=100&webp=1"
	resBody,err := utils.HttpGetRequest(url,nil,nil)
	if err != nil {
		log.Println(err.Error())
	}
	var fxhCoinInfo model.FxhCoinInfo
	err = json.Unmarshal([]byte(resBody), &fxhCoinInfo)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	for i, v := range fxhCoinInfo.Data {
		v.CDate = cDate
		fxhCoinInfo.Data[i].CDate = cDate
	}
	if len(fxhCoinInfo.Data) > 0 {
		err := global.Gdb.Save(&fxhCoinInfo.Data).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
		}
	}
}