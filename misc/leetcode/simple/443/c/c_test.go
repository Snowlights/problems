// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, longestPalindrome, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-443/problems/longest-palindrome-after-substring-concatenation-ii/
// https://leetcode.cn/problems/longest-palindrome-after-substring-concatenation-ii/
