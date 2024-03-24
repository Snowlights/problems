//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/777/C
// https://codeforces.com/problemset/status/777/problem/C
func TestCF777C(t *testing.T) {
	t.Log("Current test is [CF777C]")
	testCases := [][2]string{
		{
			`5 4
			1 2 3 5
			3 1 3 2
			4 5 2 3
			5 5 3 2
			4 4 3 4
			6
			1 1
			2 5
			4 5
			3 5
			1 3
			1 5
			`,
			`Yes
			No
			Yes
			Yes
			Yes
			No
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF777C, testCases, 0)
}
