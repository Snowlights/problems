//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1899/G
// https://codeforces.com/problemset/status/1899/problem/G
func TestCF1899G(t *testing.T) {
	t.Log("Current test is [CF1899G]")
	testCases := [][2]string{
		{
			`3
			3 5
			1 2
			2 3
			1 2 3
			1 2 2
			1 2 3
			2 3 1
			1 2 3
			2 3 3
			10 10
			2 6
			2 7
			2 4
			1 7
			2 8
			10 6
			8 5
			9 4
			3 4
			10 2 5 9 1 7 6 4 3 8
			8 9 8
			7 8 1
			7 10 6
			4 8 9
			5 5 10
			7 10 1
			9 9 2
			9 10 6
			6 6 2
			10 10 6
			1 1
			1
			1 1 1`,
			`YES
			NO
			YES
			NO
			YES
			
			NO
			YES
			YES
			YES
			NO
			YES
			YES
			NO
			NO
			NO
			
			YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1899G, testCases, 0)
}
