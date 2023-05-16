package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, distinctDifferenceArray, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}

}

// https://leetcode.cn/contest/weekly-contest-344/problems/find-the-distinct-difference-array/
// https://leetcode.cn/problems/find-the-distinct-difference-array/
