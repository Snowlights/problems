// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxSum, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-439/problems/sum-of-k-subarrays-with-length-at-least-m/
// https://leetcode.cn/problems/sum-of-k-subarrays-with-length-at-least-m/
