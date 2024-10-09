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

// https://leetcode.cn/contest/weekly-contest-406/problems/minimum-cost-for-cutting-cake-ii/
// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-ii/
