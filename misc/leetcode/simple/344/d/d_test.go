package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minIncrements, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-344/problems/make-costs-of-paths-equal-in-a-binary-tree/
// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/
