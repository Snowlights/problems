package _00_300

import (
	"fmt"
	"sort"
)

// 278
func firstBadVersion(n int) int {
	var isBadVersion func(version int) bool
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i) == true
	})
}

// 299
func getHint(secret string, guess string) string {
	right, wrong := 0, 0
	sh, gh := make(map[byte]int), make(map[byte]int)
	for i := range secret {
		if secret[i] == guess[i] {
			right++
			continue
		}
		sh[secret[i]]++
		gh[guess[i]]++
	}
	for k, v := range sh {
		wrong += min(v, gh[k])
	}

	return fmt.Sprintf("%dA%dB", right, wrong)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
