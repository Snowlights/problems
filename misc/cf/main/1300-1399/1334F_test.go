//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1334/F
// https://codeforces.com/problemset/status/1334/problem/F
func TestCF1334F(t *testing.T) {
	t.Log("Current test is [CF1334F]")
	testCases := [][2]string{
		{
			`11
			4 1 3 3 7 8 7 9 10 7 11
			3 5 0 -2 5 3 6 7 8 2 4
			3
			3 7 10
			`,
			`YES
			20`,
		},
		{
			`6
			2 1 5 3 6 5
			3 -9 0 16 22 -14
			4
			2 3 5 6
			`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1334F, testCases, 0)
}
