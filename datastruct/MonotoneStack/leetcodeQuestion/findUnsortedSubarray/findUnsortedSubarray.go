package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
最短无序连续子数组

给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

提示：

1 <= nums.length <= 104
-105 <= nums[i] <= 105

*/
func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	log.Println("最短无序连续子数组：", findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	return 0
}
