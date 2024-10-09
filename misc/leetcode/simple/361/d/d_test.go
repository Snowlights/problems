// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minOperationsQueries, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-361/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/
// https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/