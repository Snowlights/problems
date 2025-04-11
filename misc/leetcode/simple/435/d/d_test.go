package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxDifference, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-435/problems/maximum-difference-between-even-and-odd-frequency-ii/
// https://leetcode.cn/problems/maximum-difference-between-even-and-odd-frequency-ii/
