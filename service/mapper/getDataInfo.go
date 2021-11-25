package mapper

import (
	"bankroll/service/api/requestParam"
	"bankroll/service/model"
	"bankroll/utils"
)
type DataInfo struct {
}

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

//-- 查询个股信息变化
func (bankroll *DataInfo) GetStockBankroll(backrollparam requestParam.BankrollParam) ([]model.BankrollStock,error) {
	//获取时间段
	periodArr := utils.GetPeriodByOneday(backrollparam.Cdate,backrollparam.CompareNum,backrollparam.PeriodNum)
	dMap := make(map[string][]model.IndividualStock)
	for i := len(periodArr) - 1; i >= 0; i-- {
		v := periodArr[i]
		//获取数据
		individualStock, err := bankrollStockModel.GetIndividualStock(v[0],v[1],backrollparam.IndividualCode)
		if err != nil {
			return nil, err
		}
		for _, v1 := range individualStock {
			if _, ok := dMap[v1.IndividualCode]; !ok {
				dMap[v1.IndividualCode] = []model.IndividualStock{}
			}
			dMap[v1.IndividualCode] = append(dMap[v1.IndividualCode],v1)
		}
	}
	cArr := make([]model.BankrollStock, 0, 1)
	var bankrollStock = model.BankrollStock{}
	for i1, v2 := range dMap {
		if len(v2) < len(periodArr) {
			delete(dMap, i1)
			continue;
		}
		var avgRoseRatio,avgTurnoverRatio,avgRelative,avgAmplitudeRatio,avgObPrice float64
		bankrollStock.Dlist = v2
		for _, v3 := range v2 {
			avgRoseRatio, _ = utils.Add(avgRoseRatio,v3.RoseRatio)
			avgTurnoverRatio, _ = utils.Add(avgTurnoverRatio,v3.TurnoverRatio)
			avgRelative, _ = utils.Add(avgRelative,v3.Relative)
			avgAmplitudeRatio, _ = utils.Add(avgAmplitudeRatio,v3.AmplitudeRatio)
			avgObPrice, _ = utils.Add(avgObPrice,v3.ObPrice)
		}
		//平均值
		fCount := float64(len(v2));
		bankrollStock.AvgRoseRatio, _ = utils.Div(avgRoseRatio,fCount)
		bankrollStock.AvgTurnoverRatio, _ = utils.Div(avgTurnoverRatio,fCount)
		bankrollStock.AvgRelative, _ = utils.Div(avgRelative,fCount)
		bankrollStock.AvgAmplitudeRatio, _ = utils.Div(avgAmplitudeRatio,fCount)
		bankrollStock.AvgObPrice, _ = utils.Div(avgObPrice,fCount)
		//增长率
		bankrollStock.AvgObPriceRise, _ = utils.AvgRiseRatio(v2[0].ObPrice,v2[len(v2)-1].ObPrice, fCount)
		bankrollStock.CirculateStock = v2[len(v2)-1].CirculateStock
		bankrollStock.CirculateValue = v2[len(v2)-1].CirculateValue
		bankrollStock.Pe = v2[len(v2)-1].Pe
		cArr = append(cArr,bankrollStock)
	}
	return cArr,nil
}
