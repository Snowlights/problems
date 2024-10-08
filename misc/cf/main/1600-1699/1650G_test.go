//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1650/G
// https://codeforces.com/problemset/status/1650/problem/G
func TestCF1650G(t *testing.T) {
	t.Log("Current test is [CF1650G]")
	testCases := [][2]string{
		{
			`4

			4 4
			1 4
			1 2
			3 4
			2 3
			2 4
			
			6 8
			6 1
			1 4
			1 6
			1 5
			1 2
			5 6
			4 6
			6 3
			2 6
			
			5 6
			1 3
			3 5
			5 4
			3 1
			4 2
			2 1
			1 4
			
			8 18
			5 1
			2 1
			3 1
			4 2
			5 2
			6 5
			7 3
			8 4
			6 4
			8 7
			1 4
			4 7
			1 6
			6 7
			3 8
			8 5
			4 5
			4 3
			8 2`,
			`2
			4
			1
			11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1650G, testCases, 0)
}
