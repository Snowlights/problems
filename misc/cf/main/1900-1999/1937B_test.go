//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1937/B
// https://codeforces.com/problemset/status/1937/problem/B
func TestCF1937B(t *testing.T) {
	t.Log("Current test is [CF1937B]")
	testCases := [][2]string{
		{
			`3
			2
			00
			00
			4
			1101
			1100
			8
			00100111
			11101101`,
			`000
			2
			11000
			1
			001001101
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1937B, testCases, 0)
}
