// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, findXSum, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-419/problems/find-x-sum-of-all-k-long-subarrays-i/
// https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-i/
