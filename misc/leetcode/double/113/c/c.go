package main

func countPairs(coordinates [][]int, k int) int {
	cnt, ans := make(map[[2]int]int), 0
	for _, v := range coordinates {
		for i := 0; i <= k; i++ {
			ans += cnt[[2]int{v[0] ^ i, v[1] ^ (k - i)}]
		}
		cnt[[2]int{v[0], v[1]}]++
	}
	return ans
}
