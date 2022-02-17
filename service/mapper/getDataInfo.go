package mapper

import (
	"bankroll/service/api/requestParam"
	"bankroll/service/model"
	"bankroll/utils"
	"strings"
)

type DataInfo struct {
}

//-- 查询板块交易额变化
func (bankroll *DataInfo) GetPlateBankroll(backrollparam requestParam.BankrollParam) ([]model.BankrollPlate, error) {
	//获取时间段
	periodArr := utils.GetPeriodByOneday(backrollparam.Cdate, backrollparam.CompareNum, backrollparam.PeriodNum)
	dMap := make(map[string][]model.PlateBankroll)
	for i := len(periodArr) - 1; i >= 0; i-- {
		v := periodArr[i]
		//获取数据
		plateBankroll, err := bankrollPlateModel.GetPlateBankroll(v[0], v[1])
		if err != nil {
			return nil, err
		}
		for _, v1 := range plateBankroll {
			if _, ok := dMap[v1.PlateCode]; !ok {
				dMap[v1.PlateCode] = []model.PlateBankroll{}
			}
			dMap[v1.PlateCode] = append(dMap[v1.PlateCode], v1)
		}
	}
	cArr := make([]model.BankrollPlate, 0, 1)
	for i1, v2 := range dMap {
		if len(v2) < len(periodArr) {
			delete(dMap, i1)
			continue
		}
		var bankrollStock = model.BankrollPlate{}
		var avgRoseRatio, avgObVolume, avgObPrice, avgFundRealIn, avgAvgPrice, avgObPriceRise, avgObVolumeRise float64
		count := 0
		for _, v3 := range v2 {
			avgRoseRatio, _ = utils.Add(avgRoseRatio, v3.RoseRatio)
			avgObVolume, _ = utils.Add(avgObVolume, v3.ObVolume)
			avgObPrice, _ = utils.Add(avgObPrice, v3.ObPrice)
			avgFundRealIn, _ = utils.Add(avgFundRealIn, v3.FundRealIn)
			avgAvgPrice, _ = utils.Add(avgFundRealIn, v3.AvgPrice)
			//整成数组形势
			roseRatio, _ := utils.GetStructFloat64Field(v3, "RoseRatio")
			bankrollStock.RoseRatio = append(bankrollStock.RoseRatio, roseRatio)
			obVolume, _ := utils.GetStructFloat64Field(v3, "ObVolume")
			bankrollStock.ObVolume = append(bankrollStock.ObVolume, obVolume)
			obPrice, _ := utils.GetStructFloat64Field(v3, "ObPrice")
			bankrollStock.ObPrice = append(bankrollStock.ObPrice, obPrice)
			fundRealIn, _ := utils.GetStructFloat64Field(v3, "FundRealIn")
			bankrollStock.FundRealIn = append(bankrollStock.FundRealIn, fundRealIn)

			riseCompanyNum, _ := utils.GetStrucIntField(v3, "RiseCompanyNum")
			bankrollStock.RiseCompanyNum = append(bankrollStock.RiseCompanyNum, riseCompanyNum)
			dropCompanyNum, _ := utils.GetStrucIntField(v3, "DropCompanyNum")
			bankrollStock.DropCompanyNum = append(bankrollStock.DropCompanyNum, dropCompanyNum)

			avgPrice, _ := utils.GetStructFloat64Field(v3, "AvgPrice")
			bankrollStock.AvgPrice = append(bankrollStock.AvgPrice, avgPrice)

			cDate, _ := utils.GetStructStringField(v3, "CDate")
			bankrollStock.CDate = append(bankrollStock.CDate, cDate)
			//增长率
			if count != 0 {
				avgObPriceRise, _ = utils.AvgRiseRatio(v2[count-1].ObPrice, v2[count].ObPrice, 1)
				bankrollStock.AvgObPriceRise, _ = utils.Add(bankrollStock.AvgObPriceRise, avgObPriceRise)
				avgObVolumeRise, _ = utils.AvgRiseRatio(v2[count-1].ObVolume, v2[count].ObVolume, 1)
				bankrollStock.AvgObVolumeRise, _ = utils.Add(bankrollStock.AvgObVolumeRise, avgObVolumeRise)
			}
			count++
		}
		//名称
		bankrollStock.PlateCode = v2[0].PlateCode
		bankrollStock.PlateName = v2[0].PlateName
		//平均值
		fCount := float64(len(v2))
		bankrollStock.AvgRoseRatio, _ = utils.Div(avgRoseRatio, fCount)
		bankrollStock.AvgObVolume, _ = utils.Div(avgObVolume, fCount)
		bankrollStock.AvgObPrice, _ = utils.Div(avgObPrice, fCount)
		bankrollStock.AvgFundRealIn, _ = utils.Div(avgFundRealIn, fCount)
		bankrollStock.AvgAvgPrice, _ = utils.Div(avgAvgPrice, fCount)
		//增长率
		bankrollStock.AvgObPriceRise, _ = utils.AvgRiseRatio(v2[0].ObPrice, v2[len(v2)-1].ObPrice, fCount-1)
		bankrollStock.AvgObVolumeRise, _ = utils.AvgRiseRatio(v2[0].ObVolume, v2[len(v2)-1].ObVolume, fCount-1)
		cArr = append(cArr, bankrollStock)
	}
	return cArr, nil
}

//-- 查询个股信息变化
func (bankroll *DataInfo) GetStockBankroll(backrollparam requestParam.BankrollParam) ([]model.BankrollStock, error) {
	//获取时间段
	periodArr := utils.GetPeriodByOneday(backrollparam.Cdate, backrollparam.CompareNum, backrollparam.PeriodNum)
	dMap := make(map[string][]model.IndividualStock)
	for i := len(periodArr) - 1; i >= 0; i-- {
		v := periodArr[i]
		//获取数据
		individualStock, err := bankrollStockModel.GetIndividualStock(v[0], v[1], backrollparam.IndividualCode)
		if err != nil {
			return nil, err
		}
		for _, v1 := range individualStock {
			if _, ok := dMap[v1.IndividualCode]; !ok {
				dMap[v1.IndividualCode] = []model.IndividualStock{}
			}
			dMap[v1.IndividualCode] = append(dMap[v1.IndividualCode], v1)
		}
	}
	cArr := make([]model.BankrollStock, 0, 1)
	for i1, v2 := range dMap {
		if len(v2) < len(periodArr) {
			delete(dMap, i1)
			continue
		}
		var bankrollStock = model.BankrollStock{}
		var avgRoseRatio, avgTurnoverRatio, avgRelative, avgAmplitudeRatio, avgObPrice, avgObPriceRise float64
		count := 0
		for _, v3 := range v2 {
			avgRoseRatio, _ = utils.Add(avgRoseRatio, v3.RoseRatio)
			avgTurnoverRatio, _ = utils.Add(avgTurnoverRatio, v3.TurnoverRatio)
			avgRelative, _ = utils.Add(avgRelative, v3.Relative)
			avgAmplitudeRatio, _ = utils.Add(avgAmplitudeRatio, v3.AmplitudeRatio)
			avgObPrice, _ = utils.Add(avgObPrice, v3.ObPrice)
			//整成数组形势
			price, _ := utils.GetStructFloat64Field(v3, "NowPrice")
			bankrollStock.NowPrice = append(bankrollStock.NowPrice, price)
			roseRatio, _ := utils.GetStructFloat64Field(v3, "RoseRatio")
			bankrollStock.RoseRatio = append(bankrollStock.RoseRatio, roseRatio)
			turnoverRatio, _ := utils.GetStructFloat64Field(v3, "TurnoverRatio")
			bankrollStock.TurnoverRatio = append(bankrollStock.TurnoverRatio, turnoverRatio)
			relative, _ := utils.GetStructFloat64Field(v3, "Relative")
			bankrollStock.Relative = append(bankrollStock.Relative, relative)
			amplitudeRatio, _ := utils.GetStructFloat64Field(v3, "AmplitudeRatio")
			bankrollStock.AmplitudeRatio = append(bankrollStock.AmplitudeRatio, amplitudeRatio)
			obPrice, _ := utils.GetStructFloat64Field(v3, "ObPrice")
			bankrollStock.ObPrice = append(bankrollStock.ObPrice, obPrice)

			cDate, _ := utils.GetStructStringField(v3, "CDate")
			bankrollStock.CDate = append(bankrollStock.CDate, cDate)
			//增长率
			if count != 0 {
				avgObPriceRise, _ = utils.AvgRiseRatio(v2[count-1].ObPrice, v2[count].ObPrice, 1)
				bankrollStock.AvgObPriceRise, _ = utils.Add(bankrollStock.AvgObPriceRise, avgObPriceRise)
			}
			count++
		}
		//年报数据
		reportList, _ := stockReport.GetIndividualReport(i1,"")
		flag := false
		for i := len(reportList) - 1; i >= 0 ; i-- {
			nowReport := reportList[i]
			//如果是第一季度需要跳过
			hasA := strings.Contains(reportList[i].DateSort,"A")
			if (!hasA || i == len(reportList) - 1) && !flag{
				flag = true
				continue
			}
			prevReport := reportList[i + 1]
			if strings.Contains(nowReport.DateSort,"A") {
				prevReport.RetainProfit = 0
				prevReport.EarnMoney = 0
			}
			//净利润部分
			realRetainProfit, _ := utils.Sub(nowReport.RetainProfit,prevReport.RetainProfit)
			bankrollStock.RealRetainProfit = append(bankrollStock.RealRetainProfit,realRetainProfit)
			if len(bankrollStock.RealRetainProfit) > 1 {
				sumProfit := 0.00
				for _, v3 := range bankrollStock.RealRetainProfit {
					sumProfit, _ = utils.Add(v3,sumProfit)
				}
				//平均收入增长 = 本期收入 /（几个季度收入 / 几个季度 ） - 1
				flagProfit, _ := utils.Div(sumProfit, float64(len(bankrollStock.RealRetainProfit)))
				avgRise, _ := utils.Div(nowReport.RetainProfit,flagProfit)
				avgRise, _ = utils.Sub(avgRise,1)
				bankrollStock.AvgRetainRise = append(bankrollStock.AvgRetainRise,avgRise)
			}

			//rpTbRise, _ := utils.Sub(realRetainProfit,prevReport.RetainProfit)
			//rpTbRise, _ := utils.Div(rpTbRise,prevReport.RetainProfit)
			bankrollStock.RpTbRise = append(bankrollStock.RpTbRise,nowReport.RpTbRise)
			bankrollStock.RpHbRise = append(bankrollStock.RpHbRise,nowReport.RpHbRise)
			//营业收入部分
			realEarnMoney, _ := utils.Sub(nowReport.EarnMoney,prevReport.EarnMoney)
			bankrollStock.RealEarnMoney = append(bankrollStock.RealEarnMoney,realEarnMoney)
			bankrollStock.EmTbRise = append(bankrollStock.EmTbRise,nowReport.EmTbRise)
			bankrollStock.EmHbRise = append(bankrollStock.EmHbRise,nowReport.EmHbRise)

			bankrollStock.DateSort = append(bankrollStock.DateSort,nowReport.DateSort)
			bankrollStock.SellMoneyRate = append(bankrollStock.SellMoneyRate,nowReport.SellMoneyRate)
		}
		//名称
		bankrollStock.IndividualCode = v2[0].IndividualCode
		bankrollStock.IndividualName = v2[0].IndividualName
		//平均值
		fCount := float64(len(v2))
		bankrollStock.AvgRoseRatio, _ = utils.Div(avgRoseRatio, fCount)
		bankrollStock.AvgTurnoverRatio, _ = utils.Div(avgTurnoverRatio, fCount)
		bankrollStock.AvgRelative, _ = utils.Div(avgRelative, fCount)
		bankrollStock.AvgAmplitudeRatio, _ = utils.Div(avgAmplitudeRatio, fCount)
		bankrollStock.AvgObPrice, _ = utils.Div(avgObPrice, fCount)
		bankrollStock.CirculateStock = v2[len(v2)-1].CirculateStock
		bankrollStock.CirculateValue = v2[len(v2)-1].CirculateValue
		//bankrollStock.Pe = v2[len(v2)-1].Pe
		//计算ttm
		if len(bankrollStock.RealRetainProfit) >= 4 {
			ttmRealProfitArr := bankrollStock.RealRetainProfit[len(bankrollStock.RealRetainProfit) - 4:len(bankrollStock.RealRetainProfit)]
			ttmRealProfit := 0.0
			for _, f := range ttmRealProfitArr {
				ttmRealProfit, _ = utils.Add(ttmRealProfit,f)
			}
			//每股利润
			totalProfit, _ := utils.Div(ttmRealProfit,bankrollStock.CirculateStock)
			bankrollStock.Pe, _ = utils.Div(bankrollStock.NowPrice[len(bankrollStock.NowPrice)-1],totalProfit)
		}
		cArr = append(cArr, bankrollStock)
	}
	return cArr, nil
}
