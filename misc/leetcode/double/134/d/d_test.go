package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countSubarrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-134/problems/number-of-subarrays-with-and-value-of-k/
// https://leetcode.cn/problems/number-of-subarrays-with-and-value-of-k/
