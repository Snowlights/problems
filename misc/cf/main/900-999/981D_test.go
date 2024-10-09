//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/981/D
// https://codeforces.com/problemset/status/981/problem/D
func TestCF981D(t *testing.T) {
	t.Log("Current test is [CF981D]")
	testCases := [][2]string{
		{
			`10 4
			9 14 28 1 7 13 15 29 2 31`,
			`24`,
		},
		{
			`7 3
			3 14 15 92 65 35 89`,
			`64`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF981D, testCases, 0)
}
