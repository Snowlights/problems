package _00_1000

import "sort"

// 907
func sumSubarrayMins(arr []int) int {
	// 单调递增栈
	stack := []int{}
	arr = append(arr, 0)
	res := 0
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 设定prevIndex为-1 栈为空时 设定 left 左边元素个数为0
			prevIndex := -1
			if len(stack) > 0 {
				prevIndex = stack[len(stack)-1]
			}
			left := index - prevIndex - 1
			right := i - index - 1
			// 计算元素个数 * 最小值 得到res
			res += ((1 + left + right + (right * left)) * arr[index]) % 1000000007
			res = res % 1000000007
		}
		stack = append(stack, i)
	}
	return res
}

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
