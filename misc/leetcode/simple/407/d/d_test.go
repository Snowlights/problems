package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumOperations, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-407/problems/minimum-operations-to-make-array-equal-to-target/
// https://leetcode.cn/problems/minimum-operations-to-make-array-equal-to-target/
