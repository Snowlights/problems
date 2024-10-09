//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1878/F
// https://codeforces.com/problemset/status/1878/problem/F
func TestCF1878F(t *testing.T) {
	t.Log("Current test is [CF1878F]")
	testCases := [][2]string{
		{
			`7
			1 5
			1 1
			1 2
			2
			1 8
			1 9
			20 4
			1 3
			2
			1 7
			1 12
			16 10
			1 6
			1 6
			1 10
			1 9
			1 1
			1 9
			1 7
			1 3
			1 2
			1 10
			9 1
			1 3
			8 1
			1 2
			8 3
			1 5
			1 8
			1 10
			11 5
			1 8
			1 2
			1 1
			1 3
			1 1`,
			`YES
			YES
			YES
			YES
			
			YES
			NO
			YES
			
			YES
			NO
			YES
			YES
			YES
			NO
			YES
			NO
			YES
			YES
			
			NO
			
			NO
			
			YES
			NO
			NO
			
			YES
			NO
			NO
			NO
			NO
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1878F, testCases, 0)
}
