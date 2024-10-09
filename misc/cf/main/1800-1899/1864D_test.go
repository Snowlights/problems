//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1864/D
// https://codeforces.com/problemset/status/1864/problem/D
func TestCF1864D(t *testing.T) {
	t.Log("Current test is [CF1864D]")
	testCases := [][2]string{
		{
			`3
			5
			00100
			01110
			11111
			11111
			11111
			3
			100
			110
			110
			6
			010101
			111101
			011110
			000000
			111010
			001110`,
			`1
			2
			15`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1864D, testCases, 0)
}
