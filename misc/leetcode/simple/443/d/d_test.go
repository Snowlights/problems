package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minOperations, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-443/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/
// https://leetcode.cn/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/
