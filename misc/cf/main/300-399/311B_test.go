//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/311/B
// https://codeforces.com/problemset/status/311/problem/B
func TestCF311B(t *testing.T) {
	t.Log("Current test is [CF311B]")
	testCases := [][2]string{
		{
			`4 6 2
			1 3 5
			1 0
			2 1
			4 9
			1 10
			2 10
			3 12`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF311B, testCases, 0)
}
