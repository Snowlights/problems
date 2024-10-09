package _00_600

import (
	"sort"
)

// 504
func nextGreaterElements(nums []int) []int {
	ans, stack := make([]int, len(nums)), make([]int, 0)
	n := len(nums)
	for i := range ans {
		ans[i] = -1
	}
	nums = append(nums, nums...)
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top%n] = v
		}
		stack = append(stack, i)
	}

	return ans
}

// 509

func fib(n int) int {
	a, b := 0, 1
	switch n {
	case 0:
		return a
	case 1:
		return b
	default:
		for i := 2; i <= n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

// 518
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// 524
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) ||
			(len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j])
	})

	for _, v := range dictionary {
		sbegin, vbegin := 0, 0
		for sbegin < len(s) && vbegin < len(v) {
			if s[sbegin] == v[vbegin] {
				sbegin++
				vbegin++
				continue
			}
			sbegin++
		}
		if vbegin == len(v) {
			return v
		}
	}
	return ""
}
