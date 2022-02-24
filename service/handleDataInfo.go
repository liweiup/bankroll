package service

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/service/model"
	"bankroll/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HandleDataInfo struct {
	industryBankrolls []model.IndustryBankroll
	individualBankrolls []model.IndividualBankroll
	individualStocks []model.IndividualStock
	plateBankrolls []model.PlateBankroll
	relatDusDivs []model.RelatDusDiv
	stockReports []model.StockReport
}

//分发
func (hd *HandleDataInfo) HandleSwitch(fundType FundType,resBody string, sp ...interface{}) {
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(resBody))
	if err != nil {
		global.Zlog.Error(err.Error())
		return
	}
	switch fundType {
	case Conception:
		hd.handleIndustryData(fundType,doc)
	case Industry:
		hd.handleIndustryData(fundType,doc)
	case Individual:
		hd.handleIndividualData(doc)
	case IndustryStock:
		hd.handleIndividualStockData(doc,&sp)
	case Plate:
		hd.handlePlateData(doc)
	case Report:
		hd.handleReportData(doc,&sp)
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
		industryBankroll.CDate = time.Now().Format(config.DayOut)
		industryBankroll.FundType = fundTypeNum
		hd.industryBankrolls = append(hd.industryBankrolls,industryBankroll)
	})
	if len(hd.industryBankrolls) > 0 {
		err := global.Gdb.Save(&hd.industryBankrolls).Error
		if err != nil {
			global.Zlog.Error(err.Error())
		}
	}

}
//处理个股资金数据
func (hd *HandleDataInfo) handleIndividualData(doc *goquery.Document) {
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
		err := global.Gdb.Save(&hd.individualBankrolls).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
		}
	}
}

//处理个股详细数据
func (hd *HandleDataInfo) handleIndividualStockData(doc *goquery.Document,sp *[]interface{}) {
	par := *sp
	industryCode := fmt.Sprintf("%s",par[0])
	industryName := fmt.Sprintf("%s",par[1])
	cDate := time.Now().Format(config.DayOut)
	hd.individualStocks = []model.IndividualStock{}
	hd.relatDusDivs = []model.RelatDusDiv{}
	var individualStock model.IndividualStock
	var relatDusDiv model.RelatDusDiv
	relatDusDiv.IndustryCode = industryCode
	relatDusDiv.IndustryName = industryName
	relatDusDiv.CDate = cDate
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
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
		individualStock.CDate = cDate
		hd.individualStocks = append(hd.individualStocks,individualStock)
		relatDusDiv.IndividualCode = individualStock.IndividualCode
		hd.relatDusDivs = append(hd.relatDusDivs, relatDusDiv)
	})
	if len(hd.individualStocks) > 0 {
		err := global.Gdb.Save(&hd.individualStocks).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
		}
	}
	if len(hd.relatDusDivs) > 0 {
		err := global.Gdb.Save(&hd.relatDusDivs).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
		}
	}
}

//处理板块详细数据
func (hd *HandleDataInfo) handlePlateData(doc *goquery.Document) {
	hd.plateBankrolls = []model.PlateBankroll{}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var plateBankroll model.PlateBankroll
		//序号	板块	 涨跌幅(%)	总成交量（万手）	总成交额（亿元）	净流入（亿元）	上涨家数	下跌家数	均价	领涨股	最新价	涨跌幅(%)
		plateBankroll.PlateName,plateBankroll.PlateCode = getCodeAndName(s,1)
		plateBankroll.RoseRatio = utils.PercentNumToFloat(s.Children().Eq(2).Text())
		plateBankroll.ObVolume = utils.ConverMoney(s.Children().Eq(3).Text()+"万")
		plateBankroll.ObPrice = utils.ConverMoney(s.Children().Eq(4).Text()+"亿")
		plateBankroll.FundRealIn = utils.ConverMoney(s.Children().Eq(5).Text()+"亿")
		plateBankroll.RiseCompanyNum, _ = strconv.Atoi(s.Children().Eq(6).Text())
		plateBankroll.DropCompanyNum, _ = strconv.Atoi(s.Children().Eq(7).Text())
		plateBankroll.AvgPrice, _ = strconv.ParseFloat(s.Children().Eq(8).Text(),64)
		plateBankroll.CDate = time.Now().Format(config.DayOut)
		hd.plateBankrolls = append(hd.plateBankrolls,plateBankroll)
	})
	if len(hd.plateBankrolls) > 0 {
		err := global.Gdb.Save(&hd.plateBankrolls).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
		}
	}
}

//处理报告详细数据
func (hd *HandleDataInfo) handleReportData(doc *goquery.Document,sp *[]interface{}) {
	par := *sp
	industryCode := fmt.Sprintf("%s",par[0])
	hd.stockReports = []model.StockReport{}
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var stockReport model.StockReport
		stockReport.IndividualCode = industryCode
		//序号 报告期	公告日期	营业收入（元）同比增长（%）季度环比增长（%）净利润（元）同比增长（%）季度环比增长（%） 每股收益（元）每股净资产（元） 净资产收益率（%）每股经营现金流量（元）销售毛利率（%）
		stockReport.ReportDate = s.Children().Eq(1).Text()
		stockReport.AnnounceDate = s.Children().Eq(2).Text()
		stockReport.EarnMoney = utils.ConverMoney(s.Children().Eq(3).Text())
		stockReport.EmTbRise = utils.PercentNumToFloat(s.Children().Eq(4).Text() + "%")
		stockReport.EmHbRise = utils.PercentNumToFloat(s.Children().Eq(5).Text() + "%")
		stockReport.RetainProfit = utils.ConverMoney(s.Children().Eq(6).Text())
		stockReport.RpTbRise = utils.PercentNumToFloat(s.Children().Eq(7).Text() + "%")
		stockReport.RpHbRise = utils.PercentNumToFloat(s.Children().Eq(8).Text() + "%")
		stockReport.EsProfit, _ = strconv.ParseFloat(s.Children().Eq(9).Text(),64)
		stockReport.EsAssets, _ = strconv.ParseFloat(s.Children().Eq(10).Text(),64)
		stockReport.AssetsRatio = utils.PercentNumToFloat(s.Children().Eq(11).Text() + "%")
		stockReport.EsCash, _ = strconv.ParseFloat(s.Children().Eq(12).Text(),64)
		stockReport.SellMoneyRate = utils.PercentNumToFloat(s.Children().Eq(13).Text() + "%")
		stockReport.DateSort = getSortLetter(stockReport.ReportDate)
		hd.stockReports = append(hd.stockReports,stockReport)
	})
	if len(hd.stockReports) > 0 {
		err := global.Gdb.Save(&hd.stockReports).Error
		if err != nil {
			global.Zlog.Warn(err.Error())
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

func getSortLetter(s string) string {
	reportYear := string([]byte(s)[:4])
	if strings.Contains(s,"年年") {
		return reportYear+"D"
	}
	if strings.Contains(s,"三季") {
		return reportYear+"C"
	}
	if strings.Contains(s,"年中") {
		return reportYear+"B"
	}
	if strings.Contains(s,"一季") {
		return reportYear+"A"
	}
	return ""
}

//问财参数
var WenCaiParam =  map[string]string{
	"perpage": "50",
	"page": "1",
	"log_info": "{\"input_type\":\"typewrite\"}",
	"source": "Ths_iwencai_Xuangu",
	"version": "2.0",
	"query_area": "",
	"block_list": "",
	"add_info": "{\"urp\":{\"scene\":1,\"company\":1,\"business\":1},\"contentType\":\"json\",\"searchInfo\":true}",
}

//问财股票检索
func WenCaiSearch(question string,fundtype FundType) (*simplejson.Json,error) {
	url := "https://www.iwencai.com/customized/chart/get-robot-data"
	header := map[string]string{
		"hexin-v" : getHexinV(),
	}
	WenCaiParam["question"] = question
	WenCaiParam["secondary_intent"] = string(fundtype)
	resBody,err := utils.HttpPostRequestBatchorder(url,WenCaiParam,header)
	if err != nil {
		log.Println(err.Error())
		return nil,err
	}
	res, err := simplejson.NewJson([]byte(resBody))
	if err != nil {
		log.Println(err.Error())
		return nil,err
	}
	return res,nil

}
