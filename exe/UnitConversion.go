package exe

import (
	"strconv"
)

const (
	k float64 = 1e-3
	M float64 = 1e-6 
	G float64 = 1e-9
	T float64 = 1e-12
	P float64 = 1e-15
)

func DigitCount(num int) int {
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

func Conversion(num float64) string {
	digit := DigitCount(int(num))
	switch {
	case 0 < digit && digit < 4:
		return strconv.FormatFloat(num, 'f', 0, 64)
	case 3 < digit && digit < 7:
		return strconv.FormatFloat(num * k, 'f', 0, 64) + "KB"
	case 6 < digit && digit < 10:
		return strconv.FormatFloat(num * M, 'f', 0, 64) + "MB"
	case 9 < digit && digit < 13:
		return strconv.FormatFloat(num * G, 'f', 0, 64) + "GB"
	case 12< digit && digit < 16:
		return strconv.FormatFloat(num * T, 'f', 0, 64) + "TB"
	case 15 < digit:
		return strconv.FormatFloat(num * P, 'f', 0, 64) + "PB"
	default:
		return "-"
	}
}