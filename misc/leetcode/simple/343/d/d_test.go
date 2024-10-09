package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, smallestBeautifulString, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-343/problems/lexicographically-smallest-beautiful-string/
// https://leetcode.cn/problems/lexicographically-smallest-beautiful-string/
