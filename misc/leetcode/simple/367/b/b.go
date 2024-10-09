package main

import (
	"strings"
)

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans, cnt, left := s, 0, 0
	for right, b := range s {
		cnt += int(b & 1)
		for cnt > k || s[left] == '0' {
			cnt -= int(s[left] & 1)
			left++
		}
		if cnt == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}
