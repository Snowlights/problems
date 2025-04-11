//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2045/B
// https://codeforces.com/problemset/status/2045/problem/B
func TestCF2045B(t *testing.T) {
	t.Log("Current test is [CF2045B]")
	testCases := [][2]string{
		{
			`64 35 3`,
			`60`,
		},
		{
			`2024 2023 1273`,
			`1273`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF2045B, testCases, 0)
}
