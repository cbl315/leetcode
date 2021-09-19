/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
*/
package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return getMaxSubArrayByDP(nums)
}

func getMaxSubArrayByDP(nums []int) int {
	// dp expression: s[i] = maxSum(maxSuArray(nums[:i-1]), 0) + nums[i]
	maxSum := nums[0]
	dpExp := make([]int, len(nums))
	dpExp[0] = nums[0]
	for index := 1; index < len(nums); index++ {
		num := nums[index]
		if dpExp[index-1] > 0 {
			dpExp[index] = dpExp[index-1] + num
		} else {
			dpExp[index] = 0 + num
		}
		if dpExp[index] > maxSum {
			maxSum = dpExp[index]
		}
	}
	return maxSum
}
