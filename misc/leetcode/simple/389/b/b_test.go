// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countSubstrings, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-389/problems/count-substrings-starting-and-ending-with-given-character/
// https://leetcode.cn/problems/count-substrings-starting-and-ending-with-given-character/