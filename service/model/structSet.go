package model

//序号	行业	行业指数	涨跌幅	流入资金(亿)	流出资金(亿)	净额(亿)	公司家数	领涨股	涨跌幅	当前价(元)
//行业资金
type IndustryBankroll struct {
	GVA_MODEL
	FundType int `json:"fund_type"`
	DayNum int `json:"day_num"`
	IndustryCode string `json:"industry_code"`
	IndustryName string `json:"industry_name"`
	IndustryIndex     float64 `json:"industry_index"`
	RoseRatio         float64 `json:"rose_ratio"`
	FundAmountIn      float64 `json:"fund_amount_in"`
	FundAmountOut     float64 `json:"fund_amount_out"`
	FundRealIn        float64 `json:"fund_real_in"`
	CompanyNum        int     `json:"company_num"`
	LeaderCompanyName string  `json:"leader_company_name"`
	LeaderCompanyCode string  `json:"leader_company_code"`
	LeaderRoseRatio   float64 `json:"leader_rose_ratio"`
	LeaderPrice       float64 `json:"leader_price"`
	CDate             string `json:"c_date"`
}
func (IndustryBankroll) TableName() string {
	return "industry_bankroll"
}
//序号	股票代码	股票简称	最新价	涨跌幅	换手率	流入资金(元)	流出资金(元)	净额(元)	成交额(元)
//个股资金流
type IndividualBankroll struct {
	GVA_MODEL
	DayNum int `json:"day_num"`
	IndividualCode string `json:"individual_code"`
	IndividualName string `json:"individual_name"`
	EndPrice float64 `json:"end_price"`
	RoseRatio float64 `json:"rose_ratio"`
	TurnoverRatio float64 `json:"turnover_ratio"`
	FundAmountIn float64 `json:"fund_amount_in"`
	FundAmountOut float64 `json:"fund_amount_out"`
	FundRealIn float64 `json:"fund_real_in"`
	ObPrice float64 `json:"ob_price"`
	CDate  string `json:"c_date"`
}
func (IndividualBankroll) TableName() string {
	return "individual_bankroll"
}

//序号 代码 名称 现价 涨跌幅(%) 换手(%) 量比 振幅(%) 成交额 流通股 流通市值 市盈率
//个股信息
type IndividualStock struct {
	GVA_MODEL
	IndustryCode string `json:"industry_code"`
	IndividualCode string `json:"individual_code"`
	IndividualName string `json:"individual_name"`
	NowPrice float64 `json:"now_price"`
	RoseRatio float64 `json:"rose_ratio"`
	TurnoverRatio float64 `json:"turnover_ratio"`
	Relative float64 `json:"relative"`
	AmplitudeRatio float64 `json:"amplitude_ratio"`
	ObPrice float64 `json:"ob_price"`
	CirculateStock float64 `json:"circulate_stock"`
	CirculateValue float64 `json:"circulate_value"`
	Pe float64 `json:"pe"`
	CDate  string `json:"c_date"`
}
func (IndividualStock) TableName() string {
	return "individual_stock"
}