// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countPathsWithXorValue, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-146/problems/count-paths-with-the-given-xor-value/
// https://leetcode.cn/problems/count-paths-with-the-given-xor-value/
