package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximizeSumOfWeights, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-425/problems/maximize-sum-of-weights-after-edge-removals/
// https://leetcode.cn/problems/maximize-sum-of-weights-after-edge-removals/
