package model

import (
	"bankroll/global"
	"strings"
)

//个股信息
type BankrollStock struct {
	IndividualCode string `json:"individual_code" comment:"股票代码"`
	IndividualName string `json:"individual_name" comment:"股票名称"`

	ObPrice        []float64 `json:"ob_price" comment:"成交额（元）"`
	NowPrice       []float64 `json:"now_price" comment:"价格"`
	RoseRatio      []float64 `json:"rose_ratio" comment:"涨跌幅"`
	TurnoverRatio  []float64 `json:"turnover_ratio" comment:"换手率"`
	Relative       []float64 `json:"relative" comment:"量比"`
	AmplitudeRatio []float64 `json:"amplitude_ratio" comment:"振幅"`
	CDate          []string  `json:"c_date" comment:"时间"`

	RiseCompanyNum []int `json:"rise_company_num" comment:"上涨家数"`
	DropCompanyNum []int `json:"drop_company_num" comment:"下跌家数"`

	AvgRoseRatio      float64 `json:"avg_rose_ratio" comment:"平均涨跌幅"`
	AvgTurnoverRatio  float64 `json:"avg_turnover_ratio" comment:"平均换手率"`
	AvgRelative       float64 `json:"avg_relative" comment:"平均量比"`
	AvgAmplitudeRatio float64 `json:"avg_amplitude_ratio" comment:"平均振幅"`
	AvgObPrice        float64 `json:"avg_ob_price" comment:"平均成交额"`

	AvgObPriceRise float64 `json:"avg_ob_rise_price" comment:"成交量平均增长率"`
	CirculateStock float64 `json:"circulate_stock" comment:"流通股"`
	CirculateValue float64 `json:"circulate_value" comment:"流通市值（元）"`

	Pe float64 `json:"pe" comment:"市盈率"`
}

//-- 查询个股资金情况
func (bankroll *BankrollStock) GetIndividualStock(sdate, edate, individual_code string) ([]IndividualStock, error) {
	var boP = []IndividualStock{}
	r := global.Gdb.Table("individual_stock ls").Select("ls.individual_code, ls.circulate_stock, ls.individual_name, if(count(*) = 1, date_format(ls.c_date, '%m-%d'),CONCAT(min(date_format(ls.c_date, '%m-%d')),max(date_format(ls.c_date, '~%m-%d')))) AS c_date, sum(rose_ratio) AS rose_ratio, sum(turnover_ratio) AS turnover_ratio, sum(relative) /  count(*) AS relative, sum(amplitude_ratio) AS amplitude_ratio, sum(ob_price) AS ob_price, substring_index(group_concat(circulate_value),',',-1) AS circulate_value, substring_index(group_concat(now_price),',',-1) AS now_price, substring_index(group_concat(pe),',',-1) AS pe, circulate_value, count(*) AS count_num")
	r.Where("ls.c_date BETWEEN ? AND ? ", sdate, edate)
	if individual_code != "" {
		codeArr := strings.Split(individual_code, ",")
		codeStr := "'" + strings.Join(codeArr, "','") + "'"
		r.Where("individual_code in (" + codeStr + ")")
	}
	r.Group("ls.individual_code").Order("ls.individual_code asc")
	r.Scan(&boP)
	if r.Error != nil {
		return nil, r.Error
	}
	return boP, r.Error
}
