package main

func checkArray(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sumD := 0
	for i, x := range nums {
		sumD += d[i]
		x += sumD
		if x == 0 { // 无需操作
			continue
		}
		if x < 0 || i+k > n { // 无法操作
			return false
		}
		sumD -= x // 直接加到 sumD 中
		d[i+k] += x
	}
	return true
}
