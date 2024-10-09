package main

func maximumLengthSubstring(s string) (ans int) {
	cnt := [26]int{}
	left := 0
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		for cnt[b] > 2 {
			cnt[s[left]-'a']--
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
