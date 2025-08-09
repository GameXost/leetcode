package main

import (
	"fmt"
	prefix "leetcode/easy/longestPrefix"
)

func main() {
	test1 := []string{"flower", "flow", "flight"}
	fmt.Println(prefix.LongestCommonPrefix(test1))
}
