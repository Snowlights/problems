package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumDiameterAfterMerge, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-404/problems/find-minimum-diameter-after-merging-two-trees/
// https://leetcode.cn/problems/find-minimum-diameter-after-merging-two-trees/
