package main

func getWordsInLongestSubsequence(n int, words []string, groups []int) (ans []string) {
	for i, x := range groups {
		if i == n-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return
}
