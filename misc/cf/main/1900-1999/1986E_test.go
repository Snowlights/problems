//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1986/E
// https://codeforces.com/problemset/status/1986/problem/E
func TestCF1986E(t *testing.T) {
	t.Log("Current test is [CF1986E]")
	testCases := [][2]string{
		{
			`11
			1 1000000000
			1
			2 1
			624323799 708290323
			3 1
			3 2 1
			4 1
			7 1 5 3
			5 1
			11 2 15 7 10
			7 1
			1 8 2 16 8 16 31
			13 1
			2 1 1 3 3 11 12 22 45 777 777 1500 74
			10 2
			1 2 1 2 1 2 1 2 1 2
			11 2
			1 2 1 2 1 2 1 2 1 2 1
			13 3
			2 3 9 14 17 10 22 20 18 30 1 4 28
			5 1
			2 3 5 3 5`,
			`0
			83966524
			1
			4
			6
			1
			48
			-1
			0
			14
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1986E, testCases, 0)
}
