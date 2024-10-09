package main

func maximumCostSubstring(s, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
