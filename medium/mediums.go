package medium

import (
	"strconv"
	"strings"
	"unicode"
)

// func Reverse(x int) int {

// 	flag := 1
// 	if x < 0 {
// 		flag = -1
// 		x *= -1
// 	}
// 	newX := x
// 	cnt := 1
// 	for x > 9 {
// 		cnt *= 10
// 		x /= 10
// 	}
// 	res := 0
// 	for newX > 0 {
// 		res += cnt * (newX % 10)
// 		newX /= 10
// 		cnt /= 10
// 	}
// 	res *= flag
// 	if res > (1<<31)-1 || res < -(1<<31) {
// 		return 0
// 	}
// 	return res
// }

const (
	minInt32 = -1 << 31
	maxInt32 = 1<<31 - 1
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var num int32 = 0
	i := 0
	var flag int32 = 1

	switch s[i] {
	case '-':
		flag *= -1
		i++
	case '+':
		i++
	}
	for _, element := range s[i:] {
		if unicode.IsDigit(element) {
			num *= 10
			digit, _ := strconv.Atoi(string(element))
			num += int32(digit)
			if res := num * flag; res < int32(minInt32) {
				return minInt32
			} else if res > int32(maxInt32) {
				return maxInt32
			}

		} else {
			break
		}
	}

	return int(num * flag)
}
