package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, findThePrefixCommonArray, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-103/problems/find-the-prefix-common-array-of-two-arrays/
// https://leetcode.cn/problems/find-the-prefix-common-array-of-two-arrays/
