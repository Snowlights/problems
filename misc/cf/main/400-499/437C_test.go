//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/437/C
// https://codeforces.com/problemset/status/437/problem/C
func TestCF437C(t *testing.T) {
	t.Log("Current test is [CF437C]")
	testCases := [][2]string{
		{
			`4 3
			10 20 30 40
			1 4
			1 2
			2 3`,
			`40`,
		},
		{
			`4 4
			100 100 100 100
			1 2
			2 3
			2 4
			3 4`,
			`400`,
		},
		{
			`7 10
			40 10 20 10 20 80 40
			1 5
			4 7
			4 5
			5 2
			5 7
			6 4
			1 6
			1 3
			4 3
			1 4`,
			`160`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF437C, testCases, 0)
}
