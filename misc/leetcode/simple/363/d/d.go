package main

func core(n int) int {
	res := 1
	for i := 2; i*i <= n; i++ {
		e := 0
		for n%i == 0 {
			e ^= 1
			n /= i
		}
		if e == 1 {
			res *= i
		}
	}
	if n > 1 {
		res *= n
	}
	return res
}

func maximumSum(nums []int) (ans int64) {
	sum := make([]int64, len(nums)+1)
	for i, x := range nums {
		c := core(i + 1)
		sum[c] += int64(x)
		ans = max(ans, sum[c])
	}
	return
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func maximumSumV2(nums []int) (ans int64) {
	n := len(nums)
	for i := 1; i <= n; i++ {
		sum := int64(0)
		for j := 1; i*j*j <= n; j++ {
			sum += int64(nums[i*j*j-1]) // -1 是因为数组下标从 0 开始
		}
		ans = max(ans, sum)
	}
	return
}
