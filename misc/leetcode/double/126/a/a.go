package main

func sumOfEncryptedInt(nums []int) (ans int) {
	for _, x := range nums {
		mx, base := 0, 0
		for ; x > 0; x /= 10 {
			mx = max(mx, x%10)
			base = base*10 + 1
		}
		ans += mx * base
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
