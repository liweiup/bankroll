package model

import (
	"bankroll/global"
	"strings"
)

type StockReport struct {
	GVA_MODEL
	IndividualCode string  `json:"individual_code" comment:"股票编号"`
	ReportDate     string  `json:"report_date" comment:"报告期"`
	AnnounceDate   string  `json:"announce_date" comment:"公告日期"`
	DateSort       string  `json:"date_sort" comment:"排序"`
	EarnMoney      float64 `json:"earn_money" comment:"营业收入"`
	EmTbRise       float64 `json:"em_tb_rise" comment:"同比增长"`
	EmHbRise       float64 `json:"em_hb_rise" comment:"环比增长"`
	RetainProfit   float64 `json:"retain_profit" comment:"净利润"`
	RpTbRise       float64 `json:"rp_tb_rise" comment:"同比增长"`
	RpHbRise       float64 `json:"rp_hb_rise" comment:"环比增长"`
	EsProfit       float64 `json:"es_profit" comment:"每股收益"`
	EsAssets       float64 `json:"es_assets" comment:"每股净资产"`
	AssetsRatio    float64 `json:"assets_ratio" comment:"净资产收益率"`
	EsCash         float64 `json:"es_cash" comment:"每股现金流量"`
	SellMoneyRate  float64 `json:"sell_money_rate" comment:"销售毛利率"`
}
func (StockReport) TableName() string {
	return "stock_report"
}

//-- 查询报告情况
func (bankroll *StockReport) GetIndividualReport(code,report_date string) ([]StockReport, error) {
	var boP = []StockReport{}
	r := global.Gdb.Table("stock_report sr").Select("individual_code,date_sort,report_date,announce_date,earn_money,em_tb_rise,em_hb_rise,retain_profit,rp_tb_rise,rp_hb_rise,es_profit,es_assets,assets_ratio,es_cash,sell_money_rate")
	if len(code) > 0 {
		codeArr := strings.Split(code, ",")
		codeStr := "'" + strings.Join(codeArr, "','") + "'"
		r.Where("individual_code in (" + codeStr + ")")
	}
	if report_date != "" {
		r.Where("report_date like ?","%"+report_date+"%")
	}
	r.Order("date_sort desc").Limit(12)
	r.Scan(&boP)
	if r.Error != nil {
		return nil, r.Error
	}
	return boP, r.Error
}
