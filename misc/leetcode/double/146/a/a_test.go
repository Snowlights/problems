// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countSubarrays, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-146/problems/count-subarrays-of-length-three-with-a-condition/
// https://leetcode.cn/problems/count-subarrays-of-length-three-with-a-condition/
