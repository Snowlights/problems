// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumLengthSubstring, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-390/problems/maximum-length-substring-with-two-occurrences/
// https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/
