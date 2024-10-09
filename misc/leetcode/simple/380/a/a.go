package main

func maxFrequencyElements(nums []int) int {
	cnt, ans, mx := make(map[int]int), 0, 0
	for _, v := range nums {
		cnt[v]++
		c := cnt[v]
		if c > mx {
			mx = c
			ans = c
		} else if c == mx {
			ans += c
		}
	}
	return ans
}
