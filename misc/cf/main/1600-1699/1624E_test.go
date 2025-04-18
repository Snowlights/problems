//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1624/E
// https://codeforces.com/problemset/status/1624/problem/E
func TestCF1624E(t *testing.T) {
	t.Log("Current test is [CF1624E]")
	testCases := [][2]string{
		{
			`5
			
			4 8
			12340219
			20215601
			56782022
			12300678
			12345678
			
			2 3
			134
			126
			123
			
			1 4
			1210
			1221
			
			4 3
			251
			064
			859
			957
			054
			
			4 7
			7968636
			9486033
			4614224
			5454197
			9482268`,
			`3
			1 4 1
			5 6 2
			3 4 3
			-1
			2
			1 2 1
			2 3 1
			-1
			3
			1 3 2
			5 6 3
			3 4 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1624E, testCases, 0)
}
