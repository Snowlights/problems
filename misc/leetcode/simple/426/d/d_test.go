package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxTargetNodes, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-426/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/
// https://leetcode.cn/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/
