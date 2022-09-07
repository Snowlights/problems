package _00_1000

import "sort"

// 910
func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		maxv := max(nums[i]+k, nums[n-1]-k)
		minv := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, maxv-minv)
	}
	return ans
}

// 918
func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currMax, currMin := nums[0], nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currMax = max(currMax+nums[i], nums[i])
		maxSum = max(maxSum, currMax)
		currMin = min(currMin+nums[i], nums[i])
		minSum = min(minSum, currMin)
	}

	//等价于if maxSum < 0
	if total == minSum {
		return maxSum
	} else {
		return max(maxSum, total-minSum)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
