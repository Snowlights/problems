// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, smallestNumber, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-426/problems/smallest-number-with-all-set-bits/
// https://leetcode.cn/problems/smallest-number-with-all-set-bits/
