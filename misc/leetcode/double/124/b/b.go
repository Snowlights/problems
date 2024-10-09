package main

func lastNonEmptyString(s string) string {
	ans, cnt, mx := []byte{}, make(map[byte]int), 0
	for i := range s {
		cnt[s[i]]++
		if cnt[s[i]] > mx {
			mx = cnt[s[i]]
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if cnt[s[i]] == mx {
			ans = append(ans, s[i])
			cnt[s[i]] = 0
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}
