package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/gym/104363/problem/D
// https://codeforces.com/gym/104363/status/D
func TestCFD(t *testing.T) {
	t.Log("Current test is [CFD]")
	testCases := [][2]string{
		{
			`3 1 2
			`,
			`3
			`,
		},
		{
			`3 2 1
			`,
			`125
			`,
		},
		{
			`10 4 3
			`,
			`559774277
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CFD, testCases, 0)
}
