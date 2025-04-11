//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/87/D
// https://codeforces.com/problemset/status/87/problem/D
func TestCF87D(t *testing.T) {
	t.Log("Current test is [CF87D]")
	testCases := [][2]string{
		{
			`2
			2 1 5`,
			`2 1
			1`,
		},
		{
			`6
			1 2 1
			1 3 5
			3 4 2
			3 5 3
			3 6 4`,
			`16 1
			2`,
		},
		{
			`9
			6 4 72697
			9 6 72697
			1 6 38220
			2 6 38220
			6 7 72697
			6 5 72697
			8 6 72697
			3 6 38220`,
			`16 5
			1 2 5 6 7 `,
		},
		{
			`5
			5 4 58958
			2 1 37970
			2 5 37970
			1 3 37970`,
			`8 2
			1 2 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF87D, testCases, 0)
}
