package _00_1000

import (
	"sort"
	"strconv"
)

// 902
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = -1 // dp[i] = -1 表示 i 这个状态还没被计算出来
	}
	var f func(int, bool, bool) int
	f = func(i int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum { // 如果填了数字，则为 1 种合法方案
				return 1
			}
			return
		}
		if !isLimit && isNum { // 在不受到任何约束的情况下，返回记录的结果，避免重复运算
			dv := &dp[i]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 前面不填数字，那么可以跳过当前数位，也不填数字
			// isLimit 改为 false，因为没有填数字，位数都比 n 要短，自然不会受到 n 的约束
			// isNum 仍然为 false，因为没有填任何数字
			res += f(i+1, false, false)
		}
		// 根据是否受到约束，决定可以填的数字的上限
		up := byte('9')
		if isLimit {
			up = s[i]
		}
		// 注意：对于一般的题目而言，如果此时 isNum 为 false，则必须从 1 开始枚举，由于本题 digits 没有 0，所以无需处理这种情况
		for _, d := range digits { // 枚举要填入的数字 d
			if d[0] > up { // d 超过上限，由于 digits 是有序的，后面的 d 都会超过上限，故退出循环
				break
			}
			// isLimit：如果当前受到 n 的约束，且填的数字等于上限，那么后面仍然会受到 n 的约束
			// isNum 为 true，因为填了数字
			res += f(i+1, isLimit && d[0] == up, true)
		}
		return
	}
	return f(0, true, false)
}

// 904
func totalFruit(fruits []int) int {
	l, ans, h := 0, 0, make(map[int]int)
	for r, v := range fruits {
		h[v]++
		for l <= r && len(h) > 2 {
			val := fruits[l]
			h[val]--
			if h[val] == 0 {
				delete(h, val)
			}
			l++
		}
		ans = max(ans, r-l+1)
	}

	return ans
}

// 907
func sumSubarrayMins(arr []int) (ans int) {
	arr = append(arr, -1)
	st := []int{-1} // 哨兵
	for r, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			ans += arr[i] * (i - st[len(st)-1]) * (r - i) // 累加贡献
		}
		st = append(st, r)
	}
	return ans % (1e9 + 7)
}

// func sumSubarrayMins(arr []int) int {
//	n := len(arr)
//	l, r := make([]int, len(arr)), make([]int, len(arr))
//	q := []int{-1}
//	for i := range r {
//		r[i] = n
//	}
//	for i := 0; i < len(arr); i++ {
//		for len(q) > 1 && arr[i] <= arr[q[len(q)-1]] {
//			val := q[len(q)-1]
//			q = q[:len(q)-1]
//			r[val] = i
//		}
//		l[i] = q[len(q)-1]
//		q = append(q, i)
//	}
//	const mod int = 1e9+7
//	ans := 0
//	for i, v := range arr {
//		ans = ans + (i-l[i]) * (r[i] - i) * v % mod
//	}
//	return ans % mod
//}

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

// 915
func partitionDisjoint(nums []int) int {

	suf := make([]int, len(nums))
	suf[len(nums)-1] = nums[len(nums)-1]
	for j := len(nums) - 2; j >= 0; j-- {
		suf[j] = min(suf[j+1], nums[j])
	}
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if mx <= suf[i] {
			return i
		}
		mx = max(mx, nums[i])
	}
	return -1
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
