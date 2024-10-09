package main

func minimumOperations(nums []int) int {
	dp1, dp2, dp3 := 0, 0, 0
	lower := func(a, b int) bool {
		return a < b
	}
	upper := func(a, b int) bool {
		return a > b
	}
	equal := func(num, v int, eq func(a, b int) bool) int {
		if eq(num, v) {
			return 1
		}
		return 0
	}
	for _, num := range nums {
		dp1, dp2, dp3 = dp1+equal(num, 1, upper),
			min(dp1, dp2+equal(num, 2, lower))+equal(num, 2, upper),
			min(dp1, min(dp2+equal(num, 2, lower), dp3+equal(num, 3, lower)))
	}
	return dp3
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
