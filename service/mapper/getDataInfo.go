package mapper

import (
	"bankroll/global"
	"bankroll/service/model"
)

type Bankroll struct {}

var BankrollModel = new(Bankroll)

//-- 查询板块交易额
func (bankroll *Bankroll) GetPlateBankroll(sdate,edate string) ([]model.BankrollPlate,error) {
	var boP = []model.BankrollPlate{}
	r := global.GVA_DB.Raw("select ib.industry_code, ib.industry_name, sum(fund_real_in) as fund_real_in, sum(ob_price) as ob_price, count(*) as count_num from industry_bankroll ib inner join individual_stock s on ib.industry_code = s.industry_code where ib.c_date between date_sub(?, INTERVAL 0 DAY) and date_sub(?, INTERVAL 0 DAY) group by ib.industry_code order by ob_price desc",sdate,edate).Scan(&boP)
	if r.Error != nil {
		return nil,r.Error
	}
	return boP,r.Error
}
