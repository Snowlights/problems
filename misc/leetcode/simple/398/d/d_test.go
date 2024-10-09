package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, waysToReachStair, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-398/problems/find-number-of-ways-to-reach-the-k-th-stair/
// https://leetcode.cn/problems/find-number-of-ways-to-reach-the-k-th-stair/
