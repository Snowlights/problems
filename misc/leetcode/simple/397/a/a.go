package main

func findPermutationDifference(s, t string) int {
	ans, h := 0, make(map[byte]int)
	for i := range s {
		h[s[i]] = i
	}
	for i := range t {
		ans += abs(i - h[t[i]])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
