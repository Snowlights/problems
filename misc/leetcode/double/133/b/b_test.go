// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minOperations, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-133/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/
// https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/