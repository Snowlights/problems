package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxDistance, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-438/problems/maximize-the-distance-between-points-on-a-square/
// https://leetcode.cn/problems/maximize-the-distance-between-points-on-a-square/
