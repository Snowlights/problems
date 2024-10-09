//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1702/E
// https://codeforces.com/problemset/status/1702/problem/E
func TestCF1702E(t *testing.T) {
	t.Log("Current test is [CF1702E]")
	testCases := [][2]string{
		{
			`6
			4
			1 2
			4 3
			2 1
			3 4
			6
			1 2
			4 5
			1 3
			4 6
			2 3
			5 6
			2
			1 1
			2 2
			2
			1 2
			2 1
			8
			2 1
			1 2
			4 3
			4 3
			5 6
			5 7
			8 6
			7 8
			8
			1 2
			2 1
			4 3
			5 3
			5 4
			6 7
			8 6
			7 8`,
			`YES
			NO
			NO
			YES
			YES
			NO
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1702E, testCases, 0)
}
