//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1906/F
// https://codeforces.com/problemset/status/1906/problem/F
func TestCF1906F(t *testing.T) {
	t.Log("Current test is [CF1906F]")
	testCases := [][2]string{
		{
			`2 6
			1 1 -50
			1 2 -20
			2 2 -30
			1 1 60
			1 2 40
			2 2 10
			5
			1 1 6
			2 1 6
			1 1 3
			2 1 3
			1 1 2`,
			`100
			50
			0
			0
			-20`,
		},
		{
			`5 3
			1 3 3
			2 4 -2
			3 5 3
			6
			1 1 3
			2 1 3
			3 1 3
			3 2 3
			2 2 3
			2 2 2`,
			`3
			3
			4
			3
			0
			-2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1906F, testCases, 0)
}
