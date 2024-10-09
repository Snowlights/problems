package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumSum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-403/problems/find-the-minimum-area-to-cover-all-ones-ii/
// https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-ii/
