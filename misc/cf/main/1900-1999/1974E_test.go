//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1974/E
// https://codeforces.com/problemset/status/1974/problem/E
func TestCF1974E(t *testing.T) {
	t.Log("Current test is [CF1974E]")
	testCases := [][2]string{{
		`7
			1 10
			1 5
			2 80
			0 10
			200 100
			3 100
			70 100
			100 200
			150 150
			5 8
			3 1
			5 3
			3 4
			1 5
			5 3
			2 5
			1 5
			2 1
			5 3
			2 5
			2 4
			4 1
			5 1
			3 4
			5 2
			2 1
			1 2
			3 5
			3 2
			3 2`,
		`0
			10
			200
			15
			1
			9
			9`,
	},
	}
	codeforces.AssertEqualStringCase(t, CF1974E, testCases, 0)
}
