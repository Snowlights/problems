package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, findPermutation, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-397/problems/find-the-minimum-cost-array-permutation/
// https://leetcode.cn/problems/find-the-minimum-cost-array-permutation/
