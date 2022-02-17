package model

import (
	"bankroll/global"
)

type BankrollPlate struct {
	PlateCode      string `json:"plate_code" comment:"行业编号"`
	PlateName      string `json:"plate_name" comment:"行业名称"`
	RiseCompanyNum []int  `json:"rise_company_num" comment:"上涨家数"`
	DropCompanyNum []int  `json:"drop_company_num" comment:"下跌家数"`

	RoseRatio  []float64 `json:"rose_ratio" comment:"涨跌幅"`
	ObVolume   []float64 `json:"ob_volume" comment:"总成交量（手）"`
	ObPrice    []float64 `json:"ob_price" comment:"成交额（元）"`
	FundRealIn []float64 `json:"fund_real_in" comment:"净流入额"`
	AvgPrice   []float64 `json:"avg_price" comment:"均价"`
	CDate      []string  `json:"c_date" comment:"时间"`

	AvgRoseRatio  float64 `json:"avg_rose_ratio" comment:"平均涨跌幅"`
	AvgObVolume   float64 `json:"avg_ob_volume" comment:""`
	AvgObPrice    float64 `json:"avg_ob_price" comment:""`
	AvgFundRealIn float64 `json:"avg_fund_real_in" comment:"净流入额"`
	AvgAvgPrice   float64 `json:"avg_avg_price" comment:"均价"`

	AvgObVolumeRise float64 `json:"avg_ob_rise_volume" comment:"成交量平均增长率"`
	AvgObPriceRise  float64 `json:"avg_ob_rise_price" comment:"成交额平均增长率"`
}

//-- 查询板块交易额
func (bankroll *BankrollPlate) GetPlateBankroll(sdate, edate string) ([]PlateBankroll, error) {
	var boP = []PlateBankroll{}
	r := global.Gdb.Table("plate_bankroll pb").Select("plate_code, plate_name, sum(rose_ratio) as rose_ratio , sum(ob_volume) as ob_volume, sum(ob_price) as ob_price, sum(fund_real_in) as fund_real_in, sum(rise_company_num) as rise_company_num, sum(drop_company_num) as drop_company_num, sum(avg_price) / count(*) as avg_price, if(count(*) = 1, date_format(pb.c_date, '%m-%d'),CONCAT(min(date_format(pb.c_date, '%m-%d')), max(date_format(pb.c_date, '~%m-%d')))) AS c_date")
	r.Where("pb.c_date BETWEEN ? AND ? ", sdate, edate)
	r.Group("pb.plate_code").Order("pb.plate_code asc")
	r.Scan(&boP)
	if r.Error != nil {
		return nil, r.Error
	}
	return boP, r.Error
}

//-- 查询板块分类
func (bankroll *BankrollPlate) GetPlateGroup() ([]RelatDusDiv, error) {
	var boP = []RelatDusDiv{}
	r := global.Gdb.Raw("select industry_name,group_concat(individual_code) as individual_code from relat_dus_div group by industry_code").Scan(&boP)
	if r.Error != nil {
		return nil, r.Error
	}
	return boP, r.Error
}

//-- 获取股票code
func (bankroll *BankrollPlate) GetIndividualCode() ([]RelatDusDiv, error) {
	var boP = []RelatDusDiv{}
	r := global.Gdb.Raw("select individual_code from relat_dus_div group by individual_code").Scan(&boP)
	if r.Error != nil {
		return nil, r.Error
	}
	return boP, r.Error
}