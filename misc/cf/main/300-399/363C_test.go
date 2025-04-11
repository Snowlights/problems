//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/363/C
// https://codeforces.com/problemset/status/363/problem/C
func TestCF363C(t *testing.T) {
	t.Log("Current test is [CF363C]")
	testCases := [][2]string{
		{
			`helloo`,
			`hello`,
		},
		{
			`woooooow`,
			`woow`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF363C, testCases, 0)
}
