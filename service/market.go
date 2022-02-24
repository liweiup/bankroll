package service

import (
	"bankroll/config"
	"bankroll/global/redigo"
	"bankroll/service/api"
	"bankroll/utils"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)
type FundType string

const (
	Individual FundType = "ggzjl" //个股资金
	Conception FundType = "gnzjl" //概念
	Industry FundType = "hyzjl" //行业
	IndustryStock FundType = "ggzjld" //个股详细数据
	Plate FundType = "plate" //板块
	Report FundType = "report" //板块

	//问财
	WenCaiZhiShu FundType = "zhishu" //板块指数
	WenCaiStock FundType = "stock" //股票
)
var hd HandleDataInfo
 //获取 个股|行业|概念 资金数据
func MarketGetBankRoll(fundType FundType,field string,page,size int) *HandleDataInfo {
	//节假日跳过
	dEx, err := redigo.Dtype.Set.SisMember(config.HolidaySet,time.Now().Format(config.LayoutDate)).Bool()
	if dEx {
		return nil
	}
	strResUrl := fmt.Sprintf("http://data.10jqka.com.cn/funds/%s/field/%s/order/desc/ajax/1/free/1/page/%d/size/%d",fundType,field,page,size)
	header := map[string]string{
		"hexin-v" : getHexinV(),
	}
	log.Printf("行业资金数据临时hexin-v：%s  url: %s",header["hexin-v"],strResUrl)
	resBody,err := utils.HttpGetRequest(strResUrl,nil,header)
	hd.HandleSwitch(fundType,resBody);
	for _,v := range hd.industryBankrolls {
		if v.IndustryCode != "" && (fundType == Industry || fundType == Conception) {
			ftype := "thshy"
			if fundType == Conception {
				ftype = "gn" //概念
			}
			MarketGetIndividualStock(ftype,"19",v.IndustryCode,v.IndustryName,1,200)
		}
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &hd
}
//获取个股详细数据
func MarketGetIndividualStock(ft,field,code,name string,page,size int) {
	strResUrl := fmt.Sprintf("http://q.10jqka.com.cn/%s/detail/field/%s/code/%s/order/desc/ajax/1/page/%d/size/%d",ft,field,code,page,size)
	header := map[string]string{
		"hexin-v" : getHexinV(),
	}
	//log.Printf("个股数据临时hexin-v：%s  url: %s",header["hexin-v"],strResUrl)
	resBody,err := utils.HttpGetRequest(strResUrl,nil,header)
	hd.HandleSwitch(IndustryStock,resBody,code,name);
	if err != nil {
		log.Println(err.Error())
	}
}
//获取板块详细数据
func MarketGetPlateBankroll(size int) {
	//节假日跳过
	dEx, err := redigo.Dtype.Set.SisMember(config.HolidaySet,time.Now().Format(config.LayoutDate)).Bool()
	if dEx {
		return
	}
	strResUrl := fmt.Sprintf("https://q.10jqka.com.cn/thshy/index/ajax/1/size/%d",size)
	header := map[string]string{
		"hexin-v" : getHexinV(),
	}
	log.Printf("板块数据临时hexin-v：%s  url: %s",header["hexin-v"],strResUrl)
	resBody,err := utils.HttpGetRequest(strResUrl,nil,header)
	hd.HandleSwitch(Plate,resBody);
	if err != nil {
		log.Println(err.Error())
	}
}

//设置需要获取财报的股票code
func SetReportCodeToRedis() {
	relatDusDiv, _ := api.ModelPlate.GetIndividualCode()
	var codeArr []interface{}
	for _, v := range relatDusDiv {
		codeArr = append(codeArr, v.IndividualCode)
	}
	redigo.Dtype.Set.SAdd(config.StockReportCode,codeArr)
	//有效期30天
	redigo.Dtype.Key.Expire(config.StockReportCode,86400 * 30)
}

//获取财报数据
func MarketGetStockReport() {
	for i := 0; i < 20; i++ {
		code, err := redigo.Dtype.Set.SPop(config.StockReportCode).String()
		if err != nil {
			return
		}
		strResUrl := fmt.Sprintf("http://data.10jqka.com.cn/ajax/yjgg/op/code/code/%s/ajax/1/free/1/",code)
		header := map[string]string{
			"hexin-v" : getHexinV(),
		}
		log.Printf("财报数据临时hexin-v：%s  url: %s",header["hexin-v"],strResUrl)
		resBody,err := utils.HttpGetRequest(strResUrl,nil,header)
		hd.HandleSwitch(Report,resBody,code);
		if err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Second * 2)
	}
}

func WenSearchBiddingData(plateQues,stockQues string) {
	res, err := WenCaiSearch(plateQues,WenCaiZhiShu)
	if err != nil {
		return 
	}
	searchDatas := res.Get("data").Get("answer").GetIndex(0).Get("txt").GetIndex(0).Get("content").Get("components").GetIndex(0).Get("data").Get("datas")
	plateArr := []string{}
	for _, v := range searchDatas.MustArray() {
		plateArr = append(plateArr, v.(map[string]interface{})["指数简称"].(string))
	}
	plateStrSearch := strings.Join(plateArr,"或")
	//查询股票的条件
	stockQues += plateStrSearch
	stockRes, err := WenCaiSearch(stockQues,WenCaiStock)
	stockResSearchDatas := stockRes.Get("data").Get("answer").GetIndex(0).Get("txt").GetIndex(0).Get("content").Get("components").GetIndex(0).Get("data").Get("datas")
	sdate := strings.Replace(time.Now().Format(config.DayOut),"-","",-1)
	log.Println("问题："+stockQues)
	stockMapArr := []map[string]string{}
	for _, v := range stockResSearchDatas.MustArray() {
		stockMap := map[string]string{}
		vmap := v.(map[string]interface{})
		stockMap["a股票简称"] = vmap["股票简称"].(string)
		stockMap["b所属同花顺行业"] = vmap["所属同花顺行业"].(string)
		stockMap["c所属概念"] = vmap["所属概念"].(string)
		stockMap["d涨跌幅:前复权"] = vmap["涨跌幅:前复权[" + sdate + "]"].(string)
		stockMap["e量比"] = vmap["量比[" + sdate + "]"].(string)
		//stockMap["f委比"] = vmap["委比[" + sdate + "]"].(string)
		stockMap["gmacd(dea值)"] = vmap["macd(dea值)[" + sdate + "]"].(string)
		stockMap["h上市天数"] = string(vmap["上市天数["+sdate+"]"].(json.Number))
		stockMap["i市盈率(pe)"] = vmap["市盈率(pe)[" + sdate + "]"].(string)
		stockMap["j股市值(不含限售股)"] = vmap["a股市值(不含限售股)[" + sdate + "]"].(string)
		stockMapArr = append(stockMapArr, stockMap)
		volPercent, _ := strconv.ParseFloat(stockMap["e量比"],64)
		//总分100分，量比和委比分别50分
		//假设第一名量比为10,得分为50
		//剩余得分为
		fmt.Println(stockMap)
		fmt.Println("\n")
		fmt.Println(fmt.Sprintf("%.4f", volPercent))
		fmt.Println(50 * math.Log(volPercent))
		fmt.Println(50 * math.Log(10))
		//50 * math.Round(volPercent)
	}

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
