package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxSubarrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-440/problems/maximize-subarrays-after-removing-one-conflicting-pair/
// https://leetcode.cn/problems/maximize-subarrays-after-removing-one-conflicting-pair/
