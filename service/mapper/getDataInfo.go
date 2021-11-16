package mapper

import (
	"bankroll/service/api/requestParam"
	"bankroll/service/model"
	"bankroll/utils"
)

type DataInfo struct {}

//-- 查询板块交易额变化
func (bankroll *DataInfo) GetPlateBankroll(backrollparam requestParam.BankrollParam) (map[string][]model.BankrollPlate,error) {
	//获取时间段
	periodArr := utils.GetPeriodByOneday(backrollparam.Cdate,backrollparam.CompareNum,backrollparam.PeriodNum)
	dMap := make(map[string][]model.BankrollPlate)
	for _, v := range periodArr {
		//获取数据
		plateBankroll, err := bankrollPlateModel.GetPlateBankroll(v[0],v[1],backrollparam.FundType)
		if err != nil {
			return nil, err
		}
		for _, v1 := range plateBankroll {
			if _, ok := dMap[v1.IndustryCode]; !ok {
				dMap[v1.IndustryCode] = []model.BankrollPlate{}
			}
			dMap[v1.IndustryCode] = append(dMap[v1.IndustryCode],v1)
		}
	}
	return dMap,nil
}
