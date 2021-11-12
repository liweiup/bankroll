package service

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/service/model"
	"bankroll/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HandleDataInfo struct {
	industryBankrolls []model.IndustryBankroll
	individualBankrolls []model.IndividualBankroll
	individualStocks []model.IndividualStock
}

//分发
func (hd *HandleDataInfo) HandleSwitch(fundType FundType,resBody string, sp ...interface{}) {
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(resBody))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	switch fundType {
	case Conception:
		hd.handleIndustryData(fundType,doc)
	case Industry:
		hd.handleIndustryData(fundType,doc)
	case Individual:
		hd.handleIndividualData(fundType,doc)
	case IndustryStock:
		hd.handleIndividualStockData(fundType,doc,&sp)
	default:

	}
}
//处理 行业｜概念数据
func (hd *HandleDataInfo) handleIndustryData(fundType FundType,doc *goquery.Document) {
	hd.industryBankrolls = []model.IndustryBankroll{}
	fundTypeNum := 2
	if strings.EqualFold(string(fundType), string(Conception)) {
		fundTypeNum = 1
	}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var industryBankroll model.IndustryBankroll
		//序号 行业	行业指数	涨跌幅	流入资金(亿)	流出资金(亿)	净额(亿)	公司家数	领涨股	涨跌幅	当前价(元)
		industryBankroll.IndustryName,industryBankroll.IndustryCode = getCodeAndName(s,1)
		industryBankroll.IndustryIndex, _ = strconv.ParseFloat(s.Children().Eq(2).Text(),64)
		industryBankroll.RoseRatio = utils.PercentNumToFloat(s.Children().Eq(3).Text())
		industryBankroll.FundAmountIn,_ = strconv.ParseFloat(s.Children().Eq(4).Text(),64)
		industryBankroll.FundAmountIn *= float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.FundAmountOut,_ = strconv.ParseFloat(s.Children().Eq(5).Text(),64)
		industryBankroll.FundAmountOut *= float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.FundRealIn,_ = strconv.ParseFloat(s.Children().Eq(6).Text(),64)
		industryBankroll.FundRealIn *= float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.CompanyNum, _ = strconv.Atoi(s.Children().Eq(7).Text())
		industryBankroll.LeaderCompanyName,industryBankroll.LeaderCompanyCode = getCodeAndName(s,8)
		industryBankroll.LeaderRoseRatio = utils.PercentNumToFloat(s.Children().Eq(9).Text())
		industryBankroll.LeaderPrice, _ = strconv.ParseFloat(s.Children().Eq(10).Text(),64)
		industryBankroll.CDate = time.Now().Format(config.DayOut)
		industryBankroll.FundType = fundTypeNum
		hd.industryBankrolls = append(hd.industryBankrolls,industryBankroll)
	})
	if len(hd.industryBankrolls) > 0 {
		err := global.GVA_DB.Save(&hd.industryBankrolls).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
	}

}
//处理个股资金数据
func (hd *HandleDataInfo) handleIndividualData(fundType FundType,doc *goquery.Document) {
	hd.individualBankrolls = []model.IndividualBankroll{}

	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var individualBankroll model.IndividualBankroll
		//序号	股票代码	股票简称	最新价	涨跌幅	换手率	流入资金(元)	流出资金(元)	净额(元)	成交额(元)
		individualBankroll.IndividualName,individualBankroll.IndividualCode = getCodeAndName(s,2)
		individualBankroll.EndPrice,_ = strconv.ParseFloat(s.Children().Eq(3).Text(),64)
		individualBankroll.RoseRatio = utils.PercentNumToFloat(s.Children().Eq(4).Text())
		individualBankroll.TurnoverRatio = utils.PercentNumToFloat(s.Children().Eq(5).Text())
		individualBankroll.FundAmountIn = utils.ConverMoney(s.Children().Eq(6).Text())
		individualBankroll.FundAmountOut = utils.ConverMoney(s.Children().Eq(7).Text())
		individualBankroll.FundRealIn = utils.ConverMoney(s.Children().Eq(8).Text())
		individualBankroll.ObPrice = utils.ConverMoney(s.Children().Eq(9).Text())
		individualBankroll.CDate = time.Now().Format(config.DayOut)
		hd.individualBankrolls = append(hd.individualBankrolls,individualBankroll)
	})
	if len(hd.individualBankrolls) > 0 {
		err := global.GVA_DB.Save(&hd.individualBankrolls).Error
		if err != nil {
			global.GVA_LOG.Warn(err.Error())
		}
	}
}

//处理个股详细数据
func (hd *HandleDataInfo) handleIndividualStockData(fundType FundType,doc *goquery.Document,sp *[]interface{}) {
	par := *sp
	code := fmt.Sprintf("%s",par[0])
	hd.individualStocks = []model.IndividualStock{}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var individualStock model.IndividualStock
		//序号	代码	名称	现价 	涨跌幅(%) 	涨跌 	涨速(%) 	换手(%) 	量比 	振幅(%) 	成交额 	流通股 	流通市值 	市盈率 	加自选
		individualStock.IndividualName,individualStock.IndividualCode = getCodeAndName(s,2)
		if s.Children().Eq(3).Text() == "--" {
			return
		}
		individualStock.NowPrice,_ = strconv.ParseFloat(s.Children().Eq(3).Text(),64)
		individualStock.RoseRatio = utils.PercentNumToFloat(s.Children().Eq(4).Text())
		individualStock.TurnoverRatio = utils.PercentNumToFloat(s.Children().Eq(7).Text())
		individualStock.Relative,_ = strconv.ParseFloat(s.Children().Eq(8).Text(),64)
		individualStock.AmplitudeRatio = utils.PercentNumToFloat(s.Children().Eq(9).Text())
		individualStock.ObPrice = utils.ConverMoney(s.Children().Eq(10).Text())
		individualStock.CirculateStock = utils.ConverMoney(s.Children().Eq(11).Text())
		individualStock.CirculateValue = utils.ConverMoney(s.Children().Eq(12).Text())
		if s.Children().Eq(13).Text() != "--" {
			individualStock.Pe,_ = strconv.ParseFloat(s.Children().Eq(13).Text(),64)
		}
		//所属于版块code
		individualStock.IndustryCode = code
		individualStock.CDate = time.Now().Format(config.DayOut)
		hd.individualStocks = append(hd.individualStocks,individualStock)
	})
	if len(hd.individualStocks) > 0 {
		err := global.GVA_DB.Save(&hd.individualStocks).Error
		if err != nil {
			global.GVA_LOG.Warn(err.Error())
		}
	}
}

func getCodeAndName(s *goquery.Selection,eq int) (string,string){
	IndustryCodeUrl, _ := s.Children().Eq(eq).Find("a").Attr("href")
	preg,_ := regexp.Compile("[0-9]{4,10}")
	code := preg.FindString(IndustryCodeUrl)
	name := s.Children().Eq(eq).Text()
	return name,code
}