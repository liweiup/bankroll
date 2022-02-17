package utils

import (
	"math"
)

//平均增涨率
func AvgRiseRatio(f1,f2,c float64) (float64,bool) {
	if f1 == 0 {
		f1 = 1
	}
	v,b := Div(f2,f1)
	r := math.Pow(v, float64(1 / c)) - 1
	return r,b
}

func Sqrt(x float64) float64 {
	z := float64(1)
	tmp := float64(0)
	for math.Abs(tmp - z) > 0.0000000001 {
		tmp = z
		z = (z + x / z) / 2
	}
	return z
}