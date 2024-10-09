package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, colorTheArray, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-344/problems/number-of-adjacent-elements-with-the-same-color/
// https://leetcode.cn/problems/number-of-adjacent-elements-with-the-same-color/
