// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxIncreasingSubarrays, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-423/problems/adjacent-increasing-subarrays-detection-i/
// https://leetcode.cn/problems/adjacent-increasing-subarrays-detection-i/
