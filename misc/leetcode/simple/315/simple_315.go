package simple_315

// 1
func findMaxK(nums []int) int {
	h := make(map[int]bool)
	ans := -1

	for _, v := range nums {
		if h[-v] {
			ans = max(ans, abs(v))
		}
		h[v] = true
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 2

func countDistinctIntegers(nums []int) int {
	h1 := make(map[int]bool)
	for _, v := range nums {
		h1[v] = true
		tmp := 0
		for v > 0 {
			a := v % 10
			v /= 10
			tmp = tmp*10 + a
		}
		h1[tmp] = true
	}
	return len(h1)
}

// 3
func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		v, tmp := i, 0
		for v > 0 {
			a := v % 10
			v /= 10
			tmp = tmp*10 + a
		}
		if i+tmp == num {
			return true
		}
	}
	return false
}

// 4
func countSubarrays(nums []int, minK int, maxK int) int64 {
	mini, maxi, i0 := -1, -1, -1
	ans := int64(0)
	max := func(a, b int64) int64 {
		if a < b {
			return b
		}
		return a
	}

	for i, v := range nums {
		if v == minK {
			mini = i
		}
		if v == maxK {
			maxi = i
		}
		if v < minK || v > maxK {
			i0 = i
		}
		ans += max(int64(min(mini, maxi)-i0), 0)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
