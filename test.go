package main

import (
	"bankroll/global"
	"context"
)

func addDate()  {
	dateArr := make([]string,0)
	//元旦1月1日~1月3日无调休共3天
	//春节1月31日~2月6日1月29日(周六)、1月30日(周日)上班共7天
	//清明节4月3日~4月5日4月2日(周六)上班共3天
	//劳动节4月30日~5月4日4月24日(周日)、5月7日(周六)上班共5天
	//端午节6月3日~6月5日无调休共3天
	//中秋节9月10日~9月12日无调休共3天
	//国庆节10月1日~10月7日10月8日(周六)、10月9日(周日)上班共7天
	dateArr = append(dateArr, "2022-01-01")
	dateArr = append(dateArr, "2022-01-02")
	dateArr = append(dateArr, "2022-01-03")

	dateArr = append(dateArr, "2022-01-31")
	dateArr = append(dateArr, "2022-02-01")
	dateArr = append(dateArr, "2022-02-02")
	dateArr = append(dateArr, "2022-02-03")
	dateArr = append(dateArr, "2022-02-04")
	dateArr = append(dateArr, "2022-02-05")
	dateArr = append(dateArr, "2022-02-06")

	dateArr = append(dateArr, "2022-04-03")
	dateArr = append(dateArr, "2022-04-04")
	dateArr = append(dateArr, "2022-04-05")

	dateArr = append(dateArr, "2022-04-30")
	dateArr = append(dateArr, "2022-05-01")
	dateArr = append(dateArr, "2022-05-02")
	dateArr = append(dateArr, "2022-05-03")
	dateArr = append(dateArr, "2022-05-04")

	dateArr = append(dateArr, "2022-06-03")
	dateArr = append(dateArr, "2022-06-04")
	dateArr = append(dateArr, "2022-06-05")

	dateArr = append(dateArr, "2022-09-10")
	dateArr = append(dateArr, "2022-09-11")
	dateArr = append(dateArr, "2022-09-12")

	dateArr = append(dateArr, "2022-10-01")
	dateArr = append(dateArr, "2022-10-02")
	dateArr = append(dateArr, "2022-10-03")
	dateArr = append(dateArr, "2022-10-04")
	dateArr = append(dateArr, "2022-10-05")
	dateArr = append(dateArr, "2022-10-06")
	dateArr = append(dateArr, "2022-10-07")
	for _, s := range dateArr {
		global.GVA_REDIS.SAdd(context.Background(),"BK:HOLIDAY",s)
	}
}