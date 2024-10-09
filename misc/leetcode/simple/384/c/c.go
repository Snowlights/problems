package main

import (
	"math/bits"
	"sort"
)

func maxPalindromesAfterOperations(words []string) (ans int) {
	tot, mask := 0, 0
	for _, w := range words {
		tot += len(w)
		for _, c := range w {
			mask ^= 1 << (c - 'a')
		}
	}
	tot -= bits.OnesCount(uint(mask)) // 减去出现次数为奇数的字母

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for _, w := range words {
		tot -= len(w) / 2 * 2 // 长为奇数的字符串，长度要减一
		if tot < 0 {
			break
		}
		ans++
	}
	return
}
