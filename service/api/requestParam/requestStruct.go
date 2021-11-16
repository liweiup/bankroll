package requestParam

type BankrollParam struct {
	Sdate      string `json:"sdate"`
	Edate      string `json:"edate"`
	Cdate      string `json:"cdate"`
	FundType int `json:"fundType"`
	CompareNum int `json:"compareNum"`
	PeriodNum  int `json:"periodNum"`
}