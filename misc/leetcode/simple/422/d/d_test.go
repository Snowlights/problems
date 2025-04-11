package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countBalancedPermutations, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-422/problems/count-number-of-balanced-permutations/
// https://leetcode.cn/problems/count-number-of-balanced-permutations/
