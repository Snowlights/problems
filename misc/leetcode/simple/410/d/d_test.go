package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countOfPairs, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-410/problems/find-the-count-of-monotonic-pairs-ii/
// https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii/
