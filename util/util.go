package util

import "math"

func RoundToDec(val float64) float64 {
	return math.Round(val*100.0) / 100.0
}
