//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1761/E
// https://codeforces.com/problemset/status/1761/problem/E
func TestCF1761E(t *testing.T) {
	t.Log("Current test is [CF1761E]")
	testCases := [][2]string{
		{
			`4
			3
			011
			100
			100
			3
			000
			001
			010
			4
			0100
			1000
			0001
			0010
			6
			001100
			000011
			100100
			101000
			010001
			010010`,
			`0
			1
			1
			2
			3 4 
			3
			2 5 6 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1761E, testCases, 0)
}
