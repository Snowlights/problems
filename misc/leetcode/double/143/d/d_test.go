package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, smallestNumber, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-143/problems/smallest-divisible-digit-product-ii/
// https://leetcode.cn/problems/smallest-divisible-digit-product-ii/
