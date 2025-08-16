package easy

import "strconv"

// // Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil && list2 == nil {
// 		return nil
// 	}
// 	resList := &ListNode{Val: 0, Next: nil}
// 	header := resList
// 	for list1 != nil || list2 != nil {
// 		if list1 != nil && list2 != nil {
// 			if list1.Val <= list2.Val {
// 				resList.Val = list1.Val
// 				list1 = list1.Next
// 			} else if list1.Val > list2.Val {
// 				resList.Val = list2.Val
// 				list2 = list2.Next
// 			}
// 		} else if list1 != nil {
// 			resList.Val = list1.Val
// 			list1 = list1.Next
// 		} else if list2 != nil {
// 			resList.Val = list2.Val
// 			list2 = list2.Next
// 		}
// 		if list1 != nil || list2 != nil {
// 			resList.Next = &ListNode{Val: 0, Next: nil}
// 		}
// 		resList = resList.Next
// 	}
// 	return header
// }

// func RemoveDuplicates(nums []int) int {
// 	prev := nums[0]
// 	i := 1
// 	for i < len(nums) {
// 		if prev == nums[i] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 		} else {
// 			prev = nums[i]
// 			i++
// 		}
// 	}
// 	return len(nums)
// }

// func maximum69Number(num int) int {
// 	bytes := []byte(strconv.Itoa(num))
// 	for ind, elem := range bytes {
// 		if elem == '6' {
// 			bytes[ind] = '9'
// 			break
// 		}
// 	}
// 	res, _ := strconv.Atoi(string(bytes))
// 	return res
// }
