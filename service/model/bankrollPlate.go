package model

import (
	"bankroll/global"
)

type BankrollPlate struct {
	IndustryCode string `json:"industry_code"`
	CDate string `json:"c_date"`
	IndustryName string `json:"industry_name"`
	FundRealIn string `json:"fund_real_in"`
	ObPrice  string `json:"ob_price"`
	CirculateValue  string `json:"circulate_value"`
	ContactRatio  string `json:"contact_ratio"` //交易额占市值的百分比
	CountNum string `json:"count_num"` //统计条数
}

func (BankrollPlate) TableName() string {
	return "bankroll_plate"
}

//-- 查询板块交易额
func (bankroll *BankrollPlate) GetPlateBankroll(sdate,edate string,fundType int) ([]BankrollPlate,error) {
	var boP = []BankrollPlate{}
	r := global.Gdb.Raw("",sdate,edate,fundType).Scan(&boP)
	if r.Error != nil {
		return nil,r.Error
	}
	return boP,r.Error
}

//-- 查询板块分类
func (bankroll *BankrollPlate) GetPlateGroup() ([]RelatDusDiv,error) {
	var boP = []RelatDusDiv{}
	r := global.Gdb.Raw("select industry_name,group_concat(individual_code) as individual_code from relat_dus_div group by industry_code").Scan(&boP)
	if r.Error != nil {
		return nil,r.Error
	}
	return boP,r.Error
}



