package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countPairs, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-412/problems/count-almost-equal-pairs-ii/
// https://leetcode.cn/problems/count-almost-equal-pairs-ii/
