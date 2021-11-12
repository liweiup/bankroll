package model

type BankrollPlate struct {
	IndustryCode string `json:"industry_code"`
	IndustryName string `json:"industry_name"`
	FundRealIn string `json:"fund_real_in"`
	ObPrice  string `json:"ob_price"`
	CountNum string `json:"count_num"`
}

func (BankrollPlate) TableName() string {
	return "bankroll_plate"
}