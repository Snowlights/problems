// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, insertGreatestCommonDivisors, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-110/problems/insert-greatest-common-divisors-in-linked-list/
// https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/