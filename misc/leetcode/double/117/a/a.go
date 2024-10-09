package main

func distributeCandies(n int, limit int) int {
	var ans int
	for i := 0; i <= limit && i <= n; i++ {
		x := n - i
		if x < limit {
			ans += x + 1
		} else if x <= 2*limit {
			ans += 2*limit - x + 1
		}
	}
	return ans
}
