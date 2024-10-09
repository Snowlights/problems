package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumCost, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-392/problems/minimum-cost-walk-in-weighted-graph/
// https://leetcode.cn/problems/minimum-cost-walk-in-weighted-graph/
