// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxAdjacentDistance, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-148/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// https://leetcode.cn/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
