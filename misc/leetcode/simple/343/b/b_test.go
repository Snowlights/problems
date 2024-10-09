package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, firstCompleteIndex, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}

}

// https://leetcode.cn/contest/weekly-contest-343/problems/first-completely-painted-row-or-column/
// https://leetcode.cn/problems/first-completely-painted-row-or-column/
