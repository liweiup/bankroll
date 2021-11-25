package utils

import (
	"github.com/shopspring/decimal"
)
// 加法
func Add(f1,f2 float64) (float64,bool) {
	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)
	return d1.Add(d2).Float64()
}

// 除法
func Div(f1,f2 float64) (float64,bool) {
	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)
	return d1.Div(d2).Float64()
}
