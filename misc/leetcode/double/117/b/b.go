package main

func distributeCandies(n int, limit int) int64 {
	var ans int64
	for i := 0; i <= limit && i <= n; i++ {
		x := n - i
		if x < limit {
			ans += int64(x + 1)
		} else if x <= 2*limit {
			ans += int64(2*limit - x + 1)
		}
	}
	return ans
}
