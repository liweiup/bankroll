package service

import (
	"bankroll/config"
	"bankroll/global/redigo"
	"bankroll/service/api"
	"bankroll/utils"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"os"
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
		code = "002739"
		code = "603987"
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
