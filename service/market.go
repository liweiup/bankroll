package service

import (
	"bankroll/config"
	"bankroll/global"
	"bankroll/global/redigo"
	"bankroll/service/api"
	"bankroll/service/model"
	"bankroll/utils"
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/wxnacy/wgo/arrays"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type FundType string

const (
	Individual    FundType = "ggzjl"  //个股资金
	Conception    FundType = "gnzjl"  //概念
	Industry      FundType = "hyzjl"  //行业
	IndustryStock FundType = "ggzjld" //个股详细数据
	Plate         FundType = "plate"  //板块
	Report        FundType = "report" //板块
	//问财
	WenCaiZhiShu FundType = "zhishu" //板块指数
	WenCaiStock  FundType = "stock"  //股票
)




var hd HandleDataInfo

//获取 个股|行业|概念 资金数据
func MarketGetBankRoll(fundType FundType, field string, page, size int) *HandleDataInfo {
	//节假日跳过
	dEx, err := redigo.Dtype.Set.SisMember(config.HolidaySet, time.Now().Format(config.LayoutDate)).Bool()
	if dEx {
		return nil
	}
	strResUrl := fmt.Sprintf("http://data.10jqka.com.cn/funds/%s/field/%s/order/desc/ajax/1/free/1/page/%d/size/%d", fundType, field, page, size)
	header := map[string]string{
		"hexin-v": getHexinV(),
	}
	log.Printf("行业资金数据临时hexin-v：%s  url: %s", header["hexin-v"], strResUrl)
	resBody, err := utils.HttpGetRequest(strResUrl, nil, header)
	hd.HandleSwitch(fundType, resBody)
	for _, v := range hd.industryBankrolls {
		if v.IndustryCode != "" && (fundType == Industry || fundType == Conception) {
			ftype := "thshy"
			if fundType == Conception {
				ftype = "gn" //概念
			}
			MarketGetIndividualStock(ftype, "19", v.IndustryCode, v.IndustryName, 1, 200)
		}
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &hd
}

//获取个股详细数据
func MarketGetIndividualStock(ft, field, code, name string, page, size int) {
	strResUrl := fmt.Sprintf("http://q.10jqka.com.cn/%s/detail/field/%s/code/%s/order/desc/ajax/1/page/%d/size/%d", ft, field, code, page, size)
	header := map[string]string{
		"hexin-v": getHexinV(),
	}
	//log.Printf("个股数据临时hexin-v：%s  url: %s",header["hexin-v"],strResUrl)
	resBody, err := utils.HttpGetRequest(strResUrl, nil, header)
	hd.HandleSwitch(IndustryStock, resBody, code, name)
	if err != nil {
		log.Println(err.Error())
	}
}

//获取板块详细数据
func MarketGetPlateBankroll(size int) {
	//节假日跳过
	dEx, err := redigo.Dtype.Set.SisMember(config.HolidaySet, time.Now().Format(config.LayoutDate)).Bool()
	if dEx {
		return
	}
	strResUrl := fmt.Sprintf("https://q.10jqka.com.cn/thshy/index/ajax/1/size/%d", size)
	header := map[string]string{
		"hexin-v": getHexinV(),
	}
	log.Printf("板块数据临时hexin-v：%s  url: %s", header["hexin-v"], strResUrl)
	resBody, err := utils.HttpGetRequest(strResUrl, nil, header)
	hd.HandleSwitch(Plate, resBody)
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
	redigo.Dtype.Set.SAdd(config.StockReportCode, codeArr)
	//有效期30天
	redigo.Dtype.Key.Expire(config.StockReportCode, 86400*30)
}

//获取财报数据
func MarketGetStockReport() {
	for i := 0; i < 20; i++ {
		code, err := redigo.Dtype.Set.SPop(config.StockReportCode).String()
		if err != nil {
			return
		}
		strResUrl := fmt.Sprintf("http://data.10jqka.com.cn/ajax/yjgg/op/code/code/%s/ajax/1/free/1/", code)
		header := map[string]string{
			"hexin-v": getHexinV(),
		}
		log.Printf("财报数据临时hexin-v：%s  url: %s", header["hexin-v"], strResUrl)
		resBody, err := utils.HttpGetRequest(strResUrl, nil, header)
		hd.HandleSwitch(Report, resBody, code)
		if err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Second * 2)
	}
}

var emailFlagArr = []string{}
var plateCacheKey = "PLATE:CACHE"
func WenSearchBiddingData(plateQues, stockQues string) {
	wc := WenCai{plateQues,WenCaiZhiShu}
	sdate := strings.Replace(time.Now().Format(config.DayOut), "-", "", -1)
	token := utils.GetToken()
	emailText,plateStrSearch := "",""
	//获取缓存里的题材
	plateCacheStr, _ := redigo.Dtype.String.Get(plateCacheKey).String()
	if plateCacheStr != "" {
		plateStrSearch = plateCacheStr
	} else {
		res, err := wc.WenCaiSearch()
		if err != nil {
			global.Zlog.Info("请求出错：" +err.Error())
			return
		}
		searchDatas := res.Get("data").Get("answer").GetIndex(0).Get("txt").GetIndex(0).Get("content").Get("components").GetIndex(0).Get("data").Get("datas")
		plateArr := []string{}
		delPlateArr := []string{"壳资源","新股与次新股","ST板块","融资融券","核准制次新股","创投","摘帽"}
		for _, v := range searchDatas.MustArray() {
			zsName := v.(map[string]interface{})["指数简称"].(string)
			if arrays.ContainsString(delPlateArr,zsName) != -1 {
				continue
			}
			plateArr = append(plateArr, zsName)
		}
		if len(plateArr) < 1 {
			utils.SendMsg(token,utils.GetModelMsg("o7Plv6DecgmxbdFJKzwysnxM4_mc","xGBpZCsR9pOil-LzhEH0Q73Ul-SZjQK2M5ZO50NsN0s","","集合竞价筛股","没有板块结果",sdate));
			utils.SendEmail("集合竞价筛股", "问题："+plateQues+"\n 结果：没有板块结果")
			return
		} else {
			plateStrSearch = strings.Join(plateArr, "或")
			redigo.Dtype.String.Set(plateCacheKey,plateStrSearch,5 * 60)
		}
	}
	//查询股票的条件
	if plateStrSearch != "" {
		stockQues += "行业属于"+plateStrSearch
	}
	wc.Question = stockQues
	wc.fundtype = WenCaiStock
	stockRes, err := wc.WenCaiSearch()
	if err != nil {
		global.Zlog.Info("请求出错：" +err.Error())
		return
	}
	//rJson,_ := json.Marshal(stockRes)
	//global.Zlog.Info("返回结果：" +string(rJson))
	stockResSearchDatas := stockRes.Get("data").Get("answer").GetIndex(0).Get("txt").GetIndex(0).Get("content").Get("components").GetIndex(0).Get("data").Get("datas")
	szSdate := strings.Replace(time.Now().Add(-time.Hour * 24).Format(config.DayOut), "-", "", -1)
	global.Zlog.Info("问题：" + stockQues)
	for _, v := range stockResSearchDatas.MustArray() {
		emailText = ""
		stockMap := map[string]string{}
		vmap := v.(map[string]interface{})
		//vjson, _ := json.Marshal(vmap)
		//global.Zlog.Info(string(vjson))
		if vmap["股票简称"] != nil {
			//确保同只股票邮件只发一次
			flagNum := 0
			for _, e := range emailFlagArr {
				if e == vmap["股票简称"].(string) {
					flagNum = 1
				}
			}
			if flagNum == 1 {
				continue
			}
			emailFlagArr = append(emailFlagArr, vmap["股票简称"].(string))
			stockMap["a-股票简称"] = vmap["股票简称"].(string) + "\n"
			emailText += "a-股票简称: " + stockMap["a-股票简称"]
		} else {
			continue
		}
		if vmap["所属同花顺行业"] != nil {
			stockMap["b-所属同花顺行业"] = vmap["所属同花顺行业"].(string) + "\n"
			emailText += "b-所属同花顺行业: " + stockMap["b-所属同花顺行业"]
		}
		if vmap["所属概念"] != nil {
			stockMap["c-所属概念"] = vmap["所属概念"].(string) + "\n"
			emailText += "c-所属概念: " + stockMap["c-所属概念"]
		}
		if vmap["涨跌幅:前复权["+sdate+"]"] != nil {
			stockMap["d-涨跌幅:前复权"] = vmap["涨跌幅:前复权["+sdate+"]"].(string) + "\n"
			emailText += "d-涨跌幅:前复权: " + stockMap["d-涨跌幅:前复权"]
			if vmap["最新价"] != nil {
				newPrice, _ := strconv.ParseFloat(vmap["最新价"].(string), 64)
				stockMap["dd-最新价"] += fmt.Sprintf("%.2f",newPrice) + "  涨10%价：" + fmt.Sprintf("%.2f",newPrice * 1.1)
				emailText += "dd-最新价: " + stockMap["dd-最新价"] + "\n"
			}
		}
		wbyz := 0.00
		yzCount := 0.00
		if vmap["委比["+sdate+"]"] != nil {
			stockMap["f-委比"] = vmap["委比["+sdate+"]"].(string) + "\n"
			wbyz, _ = strconv.ParseFloat(vmap["委比["+sdate+"]"].(string), 64)
			wbyz /= 25
			stockMap["l-委比因子"] = fmt.Sprintf("%.4f", wbyz)
			emailText += "f-委比: " + stockMap["f-委比"]
			yzCount += wbyz
		}
		if vmap["macd(dea值)["+sdate+"]"] != nil {
			stockMap["g-macd(dea值)"] = vmap["macd(dea值)["+sdate+"]"].(string) + "\n"
			emailText += "g-macd(dea值): " + stockMap["g-macd(dea值)"]
		}
		if vmap["上市天数["+sdate+"]"] != nil {
			stockMap["h-上市天数"] = string(vmap["上市天数["+sdate+"]"].(json.Number) + "\n")
			emailText += "h-上市天数: " + stockMap["h-上市天数"]
		}
		if vmap["市盈率(pe)["+sdate+"]"] != nil {
			stockMap["i-市盈率(pe)"] = vmap["市盈率(pe)["+sdate+"]"].(string) + "\n"
			emailText += "i-市盈率(pe): " + stockMap["i-市盈率(pe)"]
		}
		scNum := 0.00
		if vmap["a股市值(不含限售股)["+szSdate+"]"] != nil {
			scNum, _ = strconv.ParseFloat(string(vmap["a股市值(不含限售股)["+szSdate+"]"].(json.Number)), 64)
			stockMap["j-a股市值(不含限售股)"] = utils.ConvertNumToCap(scNum) + "\n"
			emailText += "j-a股市值(不含限售股): " + stockMap["j-a股市值(不含限售股)"]
		}
		if vmap["涨跌幅:前复权["+sdate+"]"] != nil {
			zfyz, _ := strconv.ParseFloat(vmap["涨跌幅:前复权["+sdate+"]"].(string), 64)
			zfyz = (zfyz + 6) / 4
			stockMap["k-涨幅因子"] = fmt.Sprintf("%.4f", zfyz)
			yzCount += zfyz
		}
		emailText += "k-涨幅因子: " + stockMap["k-涨幅因子"] + "\n"
		emailText += "l-委比因子: " + stockMap["l-委比因子"] + "\n"
		if vmap["量比["+sdate+"]"] != nil {
			stockMap["e-量比因子"] = vmap["量比["+sdate+"]"].(string)
			emailText += "e-量比因子: " + stockMap["e-量比因子"] + "\n"
			lbyz, _ := strconv.ParseFloat(vmap["量比["+sdate+"]"].(string), 64)
			yzCount += lbyz
		}
		szyz := 0.00
		if stockMap["j-a股市值(不含限售股)"] != "" {
			szyz =  6000000000 / scNum
			stockMap["m-市值因子"] = fmt.Sprintf("%.4f",szyz)
			emailText += "m-市值因子: " + stockMap["m-市值因子"] + "\n"
		}
		//计算总的因子
		stockMap["n-总因子"] = fmt.Sprintf("%.4f",yzCount * szyz)
		emailText += "n-总因子: " + stockMap["n-总因子"]
		//o7Plv6I7NfqeRkFPTh1BD4_dWD00
		utils.SendMsg(token,utils.GetModelMsg("o7Plv6DecgmxbdFJKzwysnxM4_mc,o7Plv6I7NfqeRkFPTh1BD4_dWD00","xGBpZCsR9pOil-LzhEH0Q73Ul-SZjQK2M5ZO50NsN0s","","集合竞价筛股",emailText,sdate));
	}
	if emailText == "" {
		log.Println("集合竞价筛股没有结果")
		utils.SendMsg(token,utils.GetModelMsg("o7Plv6DecgmxbdFJKzwysnxM4_mc,o7Plv6I7NfqeRkFPTh1BD4_dWD00","xGBpZCsR9pOil-LzhEH0Q73Ul-SZjQK2M5ZO50NsN0s","","集合竞价筛股","集合竞价筛股没有结果",sdate));
	}
}

//龙虎榜数据获取
func WenSearchLongHuData(stockQues string) {
	wc := WenCai{stockQues,WenCaiStock}
	stockRes, err := wc.WenCaiSearch()
	if err != nil {
		global.Zlog.Info("请求出错：" +err.Error())
		return
	}
	//rJson,_ := json.Marshal(stockRes)
	//global.Zlog.Info("返回结果：" +string(rJson))
	stockResSearchDatas := stockRes.Get("data").Get("answer").GetIndex(0).Get("txt").GetIndex(0).Get("content").Get("components").GetIndex(0).Get("data").Get("datas")
	var thxLonghuStock model.ThxLonghuStock
	fieldNameMap := map[string]string{
		"IndividualCode" : "股票简称",
		"IndividualName" : "code",
		"UpPname" : "上市板块",
		"UpReason" : "上榜原因",
		"NowPrice" : "最新价",
		"RoseRatio" : "最新涨跌幅",
		"BuyValue" : "营业部买入金额合计",
		"SellValue" : "营业部卖出金额合计",
		"RealValue" : "营业部净额合计",
	}
	hourNum := 24
	if time.Now().Hour() > 15 {
		hourNum = 0
	}
	value,_ := wc.handleWcData(stockResSearchDatas, time.Duration(hourNum),&thxLonghuStock,fieldNameMap)
	for _, rv := range value {
		err := json.Unmarshal(rv,&thxLonghuStock)
		if err != nil {
			return
		}
		global.Gdb.Create(&thxLonghuStock)
	}
}

func getHexinV() string {
	jsFile := "./js/aes.min.js"
	bytes, err := ioutil.ReadFile(jsFile)
	if err != nil {
		log.Fatal("加载文件出错。。")
		os.Exit(300)
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	enc, err := vm.Call("get_v", nil)
	hexinV := fmt.Sprintf("%v", enc)
	return hexinV
}