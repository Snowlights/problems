// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, subarraySum, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-433/problems/sum-of-variable-length-subarrays/
// https://leetcode.cn/problems/sum-of-variable-length-subarrays/
