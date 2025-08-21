// package medium

// // // func Reverse(x int) int {

// // // 	flag := 1
// // // 	if x < 0 {
// // // 		flag = -1
// // // 		x *= -1
// // // 	}
// // // 	newX := x
// // // 	cnt := 1
// // // 	for x > 9 {
// // // 		cnt *= 10
// // // 		x /= 10
// // // 	}
// // // 	res := 0
// // // 	for newX > 0 {
// // // 		res += cnt * (newX % 10)
// // // 		newX /= 10
// // // 		cnt /= 10
// // // 	}
// // // 	res *= flag
// // // 	if res > (1<<31)-1 || res < -(1<<31) {
// // // 		return 0
// // // 	}
// // // 	return res
// // // }

// // // const (
// // // 	minInt32 = -1 << 31
// // // 	maxInt32 = 1<<31 - 1
// // // )

// // // func myAtoi(s string) int {
// // // 	s = strings.TrimSpace(s)
// // // 	if len(s) == 0 {
// // // 		return 0
// // // 	}
// // // 	var num int32 = 0
// // // 	i := 0
// // // 	var flag int32 = 1

// // // 	switch s[i] {
// // // 	case '-':
// // // 		flag *= -1
// // // 		i++
// // // 	case '+':
// // // 		i++
// // // 	}
// // // 	for _, element := range s[i:] {
// // // 		if unicode.IsDigit(element) {
// // // 			num *= 10
// // // 			digit, _ := strconv.Atoi(string(element))
// // // 			num += int32(digit)
// // // 			if res := num * flag; res < int32(minInt32) {
// // // 				return minInt32
// // // 			} else if res > int32(maxInt32) {
// // // 				return maxInt32
// // // 			}

// // // 		} else {
// // // 			break
// // // 		}
// // // 	}

// // // 	return int(num * flag)
// // // }

// // func RemoveDuplicates(nums []int) int {
// // 	i, j := 0, 1
// // 	flag := 1
// // 	if nums[len(nums)-1] == nums[len(nums)-2] {
// // 		flag = 2
// // 	}
// // 	for cnt := 0; cnt < len(nums); cnt++ {
// // 		if nums[i] == nums[j] && j-i == 2 {
// // 			copy(nums[j:], nums[j+1:])
// // 			fmt.Println(nums)
// // 		} else if nums[i] == nums[j] && j-i == 1 {
// // 			j++
// // 		} else if nums[i] != nums[j] {
// // 			i = j
// // 			j++
// // 		}
// // 	}

// // 	return i + flag
// // }

// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	k = k % n
// 	res := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		res[(i+k)%n] = nums[i]
// 	}
// 	copy(nums, res)
// }