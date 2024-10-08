package main

func longestValidSubstring(word string, forbidden []string) (ans int) {
	has := make(map[string]bool, len(forbidden))
	for _, s := range forbidden {
		has[s] = true
	}

	left := 0
	for right := range word {
		for i := right; i >= left && i > right-10; i-- {
			if has[word[i:right+1]] {
				left = i + 1 // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
				break
			}
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
