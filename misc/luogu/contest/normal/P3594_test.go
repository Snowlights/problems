//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

func TestP3369(t *testing.T) {
	t.Log("Current test is [P3369]")
	testCases := [][2]string{
		{
			`9 7 2
			3 4 1 9 4 1 7 1 3
			`,
			`5`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

