// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countKConstraintSubstrings, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-411/problems/count-substrings-that-satisfy-k-constraint-i/
// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/