// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minSum, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-369/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
// https://leetcode.cn/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
