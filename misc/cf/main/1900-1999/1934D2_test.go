//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1934/D2
// https://codeforces.com/problemset/status/1934/problem/D2
func TestCF1934D2(t *testing.T) {
	t.Log("Current test is [CF1934D2]")
	testCases := [][2]string{
		{
			`4
			1
			
			0 0
			3
			
			
			0 0
			13
			
			
			3 4
			
			0 0
			777777770001
			
			
			0 0`,
			`
			
			second
			
			
			first
			2 1
			
			
			first
			10 7
			
			1 2
			
			
			first
			777777770000 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1934D2, testCases, 0)
}
