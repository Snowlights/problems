// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, shiftDistance, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-144/problems/shift-distance-between-two-strings/
// https://leetcode.cn/problems/shift-distance-between-two-strings/
