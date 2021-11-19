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
	r := global.GVA_DB.Raw("select ib.industry_code, date_format(ib.c_date,'%m-%d') as c_date, ib.industry_name, sum(fund_real_in) as fund_real_in, sum(ob_price) as ob_price, sum(circulate_value) as circulate_value, sum(ob_price) / sum(circulate_value) as contact_ratio, count(*) as count_num from industry_bankroll ib inner join individual_stock s on ib.industry_code = s.industry_code and ib.c_date = s.c_date where ib.c_date between ? and ? and fund_type = ? group by ib.industry_code order by contact_ratio desc",sdate,edate,fundType).Scan(&boP)
	if r.Error != nil {
		return nil,r.Error
	}
	return boP,r.Error
}
