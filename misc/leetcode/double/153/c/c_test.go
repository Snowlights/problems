// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumCost, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-153/problems/minimum-cost-to-divide-array-into-subarrays/
// https://leetcode.cn/problems/minimum-cost-to-divide-array-into-subarrays/
