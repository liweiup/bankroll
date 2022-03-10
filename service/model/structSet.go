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

// ThxLonghuStock  龙湖榜个股信息
type ThxLonghuStock struct {
	GVA_MODEL
	IndividualCode string    `gorm:"column:individual_code" json:"individual_code"` //  股票编号
	IndividualName string    `gorm:"column:individual_name" json:"individual_name"` //  股票名称
	UpPname        string    `gorm:"column:up_pname" json:"up_pname"`               //  上市板块
	UpReason       string    `gorm:"column:up_reason" json:"up_reason"`             //  上榜原因
	NowPrice       float64   `gorm:"column:now_price" json:"now_price"`             //  最新价
	RoseRatio      float64   `gorm:"column:rose_ratio" json:"rose_ratio"`           //  最新涨跌幅
	BuyValue       float64   `gorm:"column:buy_value" json:"buy_value"`             //  营业部买入金额合计（元）
	SellValue      float64   `gorm:"column:sell_value" json:"sell_value"`           //  营业部卖出金额合计（元）
	RealValue      float64   `gorm:"column:real_value" json:"real_value"`           //  营业部净额合计（元）
	CDate          string `gorm:"column:c_date" json:"c_date"`
}
func (ThxLonghuStock) TableName() string {
	return "thx_longhu_stock"
}

// 非小号json结果
type FxhCoinInfo struct {
	Data []BiDealDetail `json:"data" `
}
// 币交易明细
type BiDealDetail struct {
	GVA_MODEL
	BiCode string `gorm:"column:bi_code" db:"bi_code" json:"name" form:"bi_code"` //币code
	BiName string `gorm:"column:bi_name" db:"bi_name" json:"fullname" form:"bi_name"` //币名称
	PriceUsd float64 `gorm:"column:price_usd" db:"price_usd" json:"current_price_usd" form:"price_usd"` //价格
	VolUsd float64 `gorm:"column:vol_usd" db:"vol_usd" json:"vol_usd" form:"vol_usd"` //24小时交易额
	TurnoverRatio float64 `gorm:"column:turnover_ratio" db:"turnover_ratio" json:"turnoverrate" form:"turnover_ratio"` //换手率
	RoseRatio float64 `gorm:"column:rose_ratio" db:"rose_ratio" json:"change_percent" form:"rose_ratio"` //涨跌幅
	CDate string `gorm:"column:c_date" db:"c_date" json:"c_date" form:"c_date"`
}

func (BiDealDetail) TableName() string {
	return "bi_deal_detail"
}

//微信推送模版
type WxPushModel struct {
	Touser string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL string `json:"url"`
	Data struct {
		First struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"first"`
		Keyword1 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword1"`
		Keyword2 struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"keyword2"`
		Remark struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"remark"`
	} `json:"data"`
}