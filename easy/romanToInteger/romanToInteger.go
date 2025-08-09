package romanToInteger

import "fmt"

func romanToInt(s string) int {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	flag := true
	newS := []rune(s)
	totalSum := 0
	for i := 0; i < len(s)-1; i++ {
		curVal := nums[newS[i]]
		nextVal := nums[newS[i+1]]

		// fmt.Println(curVal, nextVal, totalSum)
		if curVal >= nextVal {
			flag = true
		} else if curVal < nextVal {
			flag = false
		}
		if flag == true {
			totalSum += curVal
		} else {
			totalSum -= curVal
		}

	}
	totalSum += nums[newS[len(s)-1]]
	return totalSum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
