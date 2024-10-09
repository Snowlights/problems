package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumCost, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-343/problems/minimum-cost-of-a-path-with-special-roads/
// https://leetcode.cn/problems/minimum-cost-of-a-path-with-special-roads/
