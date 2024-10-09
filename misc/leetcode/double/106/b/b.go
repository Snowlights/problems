package main

func longestSemiRepetitiveSubstring(s string) (ans int) {
	l, cnt := 0, 0
	ans = 1
	for r := 1; r < len(s); r++ {
		if s[r] == s[r-1] {
			cnt++
			if cnt > 1 {
				l++
				for s[l] != s[l-1] {
					l++
				}
				cnt--
			}
		}
		ans = max(ans, r-l+1)
	}

	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
