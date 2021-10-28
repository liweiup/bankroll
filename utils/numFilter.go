package utils

import (
	"log"
	"strconv"
	"strings"
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