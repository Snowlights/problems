package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countKConstraintSubstrings, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-411/problems/count-substrings-that-satisfy-k-constraint-ii/
// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-ii/
