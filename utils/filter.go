package utils

import (
	"bankroll/config"
	"bankroll/global/redigo"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 百分数转小数
 */
func PercentNumToFloat(num string) float64{
	num = strings.Replace(num,"%","",1)
	floatNum, err := strconv.ParseFloat(num,64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return floatNum / 100
}

// 将十进制数字转化为二进制字符串
func ConvertToBin(num int) int {
	s := ""
	if num == 0 {
		return 0
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	rNum,_ := strconv.Atoi(s)
	return rNum
}

/**
金额转成数字
 */
func ConverMoney(m string) float64 {
	//取出数字
	preg,_ := regexp.Compile(`[-(0-9)|\.]*`)
	num, _ := strconv.ParseFloat(preg.FindString(m),64)
	if strings.LastIndex(m,"亿") != -1 {
		num *= float64(ConvertToBin(config.YiYi))
		return num
	}
	if strings.LastIndex(m,"千万") != -1 {
		num *= float64(ConvertToBin(config.QianWan))
		return num
	}
	if strings.LastIndex(m,"百万") != -1 {
		num *= float64(ConvertToBin(config.BaiWan))
		return num
	}
	if strings.LastIndex(m,"万") != -1 {
		num *= float64(ConvertToBin(config.Wan))
		return num
	}
	return 0
}
/**
获取两个日期间的所有日期
 */
func SplitDate(beginDate, endDate, format string) []string {
	bDate, _ := time.ParseInLocation(format, beginDate, time.Local)
	eDate, _ := time.ParseInLocation(format, endDate, time.Local)
	day := int(eDate.Sub(bDate).Hours() / 24)
	dlist := make([]string, 0)
	dlist = append(dlist, beginDate)
	for i := 1; i < day; i++ {
		result := bDate.AddDate(0, 0, i)
		dlist = append(dlist, result.Format(format))
	}
	if beginDate != endDate {
		dlist = append(dlist, endDate)
	}
	return dlist
}

/**
返回一个排除节假日和周末的周期
 */
func getDayFilterHoli(currentSDate,currentEDate time.Time) (sDateStr,EDateStr string) {
	currentSDateStr := currentSDate.Format(config.LayoutDate)
	//结束时间
	currentEDateStr := currentEDate.Format(config.LayoutDate)
	//间隔日期
	curlist := SplitDate(currentSDateStr, currentEDateStr, config.LayoutDate)
	cdayNum := 0
	for i,cudate := range curlist {
		//去除周六周天 节假日
		fdate, _ := time.Parse(config.LayoutDate,cudate);
		dEx, _ := redigo.Dtype.Set.SisMember(config.HolidaySet,time.Now().Format(config.LayoutDate)).Bool()
		//周六推移两天，周天和节假日推一天
		if fdate.Weekday() == time.Sunday || fdate.Weekday() == time.Saturday || dEx {
			cdayNum ++
		}
		if i == 0 && fdate.Weekday() == time.Sunday {
			cdayNum ++
		}
	}
	currentSDate = currentSDate.AddDate(0,0,-cdayNum)
	currentSDateStr =  currentSDate.Format(config.LayoutDate)
	return currentSDateStr,currentEDateStr
}
/**
根据时间返回一个周期
conpareNum 几天时间
periodNum 多少个周期
 */
func GetPeriodByOneday(dateStr string,compareNum,periodNum int) [][]string {
	periodArr := [][]string{}
	sfdate :=  ""
	for i := 0; i < periodNum; i++ {
		fdArr := []string{}
		if i== 0 {
			currentEDate,_ := time.ParseInLocation(config.LayoutDate,dateStr,time.Local);
			//当前日期天 - 统计天数 + 1 = 统计开始时间
			currentSDate := currentEDate.AddDate(0, 0, -compareNum + 1)
			currentSDateStr,currentEDateStr := getDayFilterHoli(currentSDate,currentEDate)
			sfdate = currentSDateStr
			fdArr = []string{currentSDateStr,currentEDateStr}
			periodArr = append(periodArr,fdArr)
			continue
		}
		//println(sfdate)
		//println("=================")
		currentSDate, _ := time.Parse(config.LayoutDate,sfdate)
		//上一个周期的时间
		prevSdate := currentSDate.AddDate(0,0,-compareNum)
		prevEdate := currentSDate.AddDate(0,0,-1)
		prevSDateStr,prevEDateStr := getDayFilterHoli(prevSdate,prevEdate)
		sfdate = prevSDateStr
		fdArr = []string{prevSDateStr,prevEDateStr}
		periodArr = append(periodArr,fdArr)
	}
	return periodArr
}
