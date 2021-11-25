package model
//行业资金
type IndustryBankroll struct {
	GVA_MODEL
	FundType int `json:"fund_type"  comment:"类型，1表示概念，2表示行业"`
	DayNum int `json:"day_num" comment:"几天时间"`
	IndustryCode string `json:"industry_code" comment:"行业编号"`
	IndustryName string `json:"industry_name" comment:"行业名称"`
	IndustryIndex     float64 `json:"industry_index" comment:"行业指数"`
	RoseRatio         float64 `json:"rose_ratio" comment:"涨跌幅"`
	FundAmountIn      float64 `json:"fund_amount_in" comment:"流入资金"`
	FundAmountOut     float64 `json:"fund_amount_out" comment:"流出资金"`
	FundRealIn        float64 `json:"fund_real_in" comment:"净额"`
	CompanyNum        int     `json:"company_num" comment:"公司家数"`
	CDate             string `json:"c_date" comment:"股票代码"`
}
func (IndustryBankroll) TableName() string {
	return "industry_bankroll"
}
//个股资金流
type IndividualBankroll struct {
	GVA_MODEL
	DayNum int `json:"day_num" comment:""`
	IndividualCode string `json:"individual_code" comment:"股票代码"`
	IndividualName string `json:"individual_name" comment:"股票简称"`
	EndPrice float64 `json:"end_price" comment:"最新价"`
	RoseRatio float64 `json:"rose_ratio" comment:"涨跌幅"`
	TurnoverRatio float64 `json:"turnover_ratio" comment:"换手率"`
	FundAmountIn float64 `json:"fund_amount_in" comment:"流入资金"`
	FundAmountOut float64 `json:"fund_amount_out" comment:"流出资金"`
	FundRealIn float64 `json:"fund_real_in" comment:"净额"`
	ObPrice float64 `json:"ob_price" comment:"成交额"`
	CDate  string `json:"c_date" comment:"时间"`
}
func (IndividualBankroll) TableName() string {
	return "individual_bankroll"
}

//个股信息
type IndividualStock struct {
	GVA_MODEL
	IndividualCode string `json:"individual_code" comment:"股票代码"`
	IndividualName string `json:"individual_name" comment:"股票名称"`
	NowPrice float64 `json:"now_price" comment:"现在价格"`
	RoseRatio float64 `json:"rose_ratio" comment:"涨跌幅"`
	TurnoverRatio float64 `json:"turnover_ratio" comment:"换手率"`
	Relative float64 `json:"relative" comment:"量比"`
	AmplitudeRatio float64 `json:"amplitude_ratio" comment:"振幅"`
	ObPrice float64 `json:"ob_price" comment:"成交额（元）"`
	CirculateStock float64 `json:"circulate_stock" comment:"流通股"`
	CirculateValue float64 `json:"circulate_value" comment:"流通市值（元）"`
	Pe float64 `json:"pe" comment:"市盈率"`
	CDate  string `json:"c_date" comment:"时间"`
}
func (IndividualStock) TableName() string {
	return "individual_stock"
}

//板块资金
type PlateBankroll struct {
	GVA_MODEL
	PlateCode      string  `json:"plate_code" comment:"行业编号"`
	PlateName      string  `json:"plate_name" comment:"行业名称"`
	RoseRatio      float64 `json:"rose_ratio" comment:"涨跌幅"`
	ObVolume       float64 `json:"ob_volume" comment:"总成交量（手）"`
	ObPrice        float64 `json:"ob_price" comment:"成交额（元）"`
	FundRealIn     float64 `json:"fund_real_in" comment:"净流入额"`
	RiseCompanyNum int   `json:"rise_company_num" comment:"上涨家数"`
	DropCompanyNum int   `json:"drop_company_num" comment:"下跌家数"`
	AvgPrice       float64 `json:"avg_price" comment:"均价"`
	CDate  string `json:"c_date" comment:"时间"`
}
func (PlateBankroll) TableName() string {
	return "plate_bankroll"
}

//关联关系表
type RelatDusDiv struct {
	GVA_MODEL
	IndividualCode string `json:"individual_code" comment:"股票代码"`
	IndustryCode string `json:"industry_code" comment:"行业编号"`
	IndustryName string `json:"industry_name" comment:"行业名称"`
	CDate  string `json:"c_date" comment:"时间"`
}
func (RelatDusDiv) TableName() string {
	return "relat_dus_div"
}