//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/383/C
// https://codeforces.com/problemset/status/383/problem/C
func TestCF383C(t *testing.T) {
	t.Log("Current test is [CF383C]")
	testCases := [][2]string{
		{
			`5 5
			1 2 1 1 2
			1 2
			1 3
			2 4
			2 5
			1 2 3
			1 1 2
			2 1
			2 2
			2 4
			`,
			`3
			3
			0
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF383C, testCases, 0)
}
