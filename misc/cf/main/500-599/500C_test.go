//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/500/C
// https://codeforces.com/problemset/status/500/problem/C
func TestCF500C(t *testing.T) {
	t.Log("Current test is [CF500C]")
	testCases := [][2]string{
		{
			`3 5
			1 2 3
			1 3 2 3 1
			`,
			`12`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF500C, testCases, 0)
}
