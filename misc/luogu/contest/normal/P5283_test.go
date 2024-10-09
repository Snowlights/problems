//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

func TestP5283(t *testing.T) {
	t.Log("Current test is [P5283]")
	testCases := [][2]string{
		{
			`3 2
			1 2 3
			`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
