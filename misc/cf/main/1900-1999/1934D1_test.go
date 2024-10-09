//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1934/D1
// https://codeforces.com/problemset/status/1934/problem/D1
func TestCF1934D1(t *testing.T) {
	t.Log("Current test is [CF1934D1]")
	testCases := [][2]string{
		{
			`3
			7 3
			4 2
			481885160128643072 45035996273704960`,
			`1
			7 3
			-1
			3
			481885160128643072 337769972052787200 49539595901075456 45035996273704960`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1934D1, testCases, 0)
}
