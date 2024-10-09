//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/13/E
// https://codeforces.com/problemset/status/13/problem/E
func TestCF13E(t *testing.T) {
	t.Log("Current test is [CF13E]")
	testCases := [][2]string{
		{
			`8 5
			1 1 1 1 1 2 8 2
			1 1
			0 1 3
			1 1
			0 3 4
			1 2
			`,
			`8 7
			8 5
			7 3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF13E, testCases, 0)
}
