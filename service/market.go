package service

import (
	"bankroll/config"
	"bankroll/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)
type FundType string
const (
	Individual FundType = "ggzjl" //个股
	Conception FundType = "gnzjl" //概念
	Industry FundType = "hyzjl" //行业
)
//获取行业资金数据
func MarketGetBankRoll(fundType FundType,page,size int) []IndustryBankroll {
	//概念资金
	strResUrl := fmt.Sprintf("http://data.10jqka.com.cn/funds/%s/field/tradezdf/order/desc/ajax/1/free/1/page/%d/size/%d",fundType,page,size)
	header := map[string]string{
		"hexin-v" : getHexinV(),
	}
	println(header["hexin-v"])
	resBody,err := utils.HttpGetRequest(strResUrl,nil,header)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resBody))
	if err != nil {
		log.Fatal(err.Error())
	}
	var industryBankrolls = []IndustryBankroll{}
	// Find the review items
	doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var industryBankroll IndustryBankroll
		//股票代码	股票简称	最新价	涨跌幅	换手率	流入资金(元)	流出资金(元)	净额(元)	成交额(元)

		//序号 行业	行业指数	涨跌幅	流入资金(亿)	流出资金(亿)	净额(亿)	公司家数	领涨股	涨跌幅	当前价(元)
		industryBankroll.IndustryName,industryBankroll.IndustryCode = getCodeAndName(s,1)
		industryBankroll.IndustryIndex, _ = strconv.ParseFloat(s.Children().Eq(2).Text(),64)
		industryBankroll.RoseRatio = utils.PercentNumToFloat(s.Children().Eq(3).Text())
		industryBankroll.FundAmountIn,_ = strconv.ParseFloat(s.Children().Eq(4).Text(),64)
		industryBankroll.FundAmountIn = float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.FundAmountOut,_ = strconv.ParseFloat(s.Children().Eq(5).Text(),64)
		industryBankroll.FundAmountOut *= float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.FundRealIn,_ = strconv.ParseFloat(s.Children().Eq(6).Text(),64)
		industryBankroll.FundRealIn *= float64(utils.ConvertToBin(config.YiYi))
		industryBankroll.CompanyNum, _ = strconv.Atoi(s.Children().Eq(7).Text())
		industryBankroll.LeaderCompanyName,industryBankroll.LeaderCompanyCode = getCodeAndName(s,8)
		industryBankroll.LeaderRoseRatio = utils.PercentNumToFloat(s.Children().Eq(9).Text())
		industryBankroll.LeaderPrice, _ = strconv.ParseFloat(s.Children().Eq(10).Text(),64)
		industryBankroll.CDate = time.Now().Format(config.DayOut)
		//fmt.Println(industryBankroll.IndustryName,
		//	industryBankroll.IndustryCode,
		//	industryBankroll.IndustryIndex,
		//	industryBankroll.RoseRatio,
		//	industryBankroll.FundAmountIn,
		//	industryBankroll.FundAmountOut,
		//	industryBankroll.FundRealIn,
		//	industryBankroll.CompanyNum,
		//	industryBankroll.LeaderCompanyName,industryBankroll.LeaderCompanyCode,
		//	industryBankroll.LeaderRoseRatio,
		//	industryBankroll.LeaderPrice,
		//)
		industryBankrolls = append(industryBankrolls, industryBankroll)
	})
	fmt.Println(industryBankrolls)

	return industryBankrolls
	//fmt.Println(header["hexin-v"])
}

func getCodeAndName(s *goquery.Selection,eq int) (string,int){
	IndustryCodeUrl, _ := s.Children().Eq(eq).Find("a").Attr("href")
	preg,_ := regexp.Compile("[0-9]{4,10}")
	code, _ := strconv.Atoi(preg.FindString(IndustryCodeUrl))
	name := s.Children().Eq(eq).Text()
	return name,code
}

func getHexinV() string{
	jsFile := "./js/aes.min.js"
	bytes, err := ioutil.ReadFile(jsFile)
	if err != nil {
		log.Fatal("加载文件出错。。")
		os.Exit(300)
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	enc,err := vm.Call("get_v",nil)
	hexinV := fmt.Sprintf("%v", enc)
	return hexinV
}
