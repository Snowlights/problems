package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, gcdValues, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-418/problems/sorted-gcd-pair-queries/
// https://leetcode.cn/problems/sorted-gcd-pair-queries/
