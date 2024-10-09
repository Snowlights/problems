//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1851/G
// https://codeforces.com/problemset/status/1851/problem/G
func TestCF1851G(t *testing.T) {
	t.Log("Current test is [CF1851G]")
	testCases := [][2]string{
		{
			`2
			7 7
			1 5 3 4 2 4 1
			1 4
			4 3
			3 6
			3 2
			2 5
			5 6
			5 7
			5
			1 1 3
			6 2 0
			4 7 0
			1 7 4
			1 7 2
			6 5
			4 7 6 2 5 1
			1 3
			5 3
			1 5
			2 4
			6 2
			5
			1 5 1
			1 3 1
			1 2 1000
			6 2 6
			6 2 5`,
			`YES
			NO
			YES
			YES
			NO
			
			YES
			NO
			NO
			YES
			NO`,
		},
		{
			`2
			3 2
			1 3 9
			1 2
			2 3
			5
			1 1 1
			3 2 2
			1 1 2
			3 3 0
			1 2 1
			3 3
			1 4 1
			1 2
			2 3
			1 3
			5
			3 3 9
			1 3 6
			1 1 2
			3 3 6
			3 3 4
			`,
			`YES
			YES
			YES
			YES
			NO
			
			YES
			YES
			YES
			YES
			YES
			`,
		},
		{
			`1
			6 10
			7 9 2 10 8 6
			4 2
			6 1
			4 5
			3 5
			6 4
			1 3
			2 6
			6 5
			1 2
			3 6
			5
			4 4 8
			3 3 1
			5 5 9
			2 1 7
			6 6 10
			`,
			`YES
			YES
			YES
			YES
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1851G, testCases, 0)
}
