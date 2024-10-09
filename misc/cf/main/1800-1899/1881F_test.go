//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1881/F
// https://codeforces.com/problemset/status/1881/problem/F
func TestCF1881F(t *testing.T) {
	t.Log("Current test is [CF1881F]")
	testCases := [][2]string{
		{
			`6
			7 3
			2 6 7
			1 2
			1 3
			2 4
			2 5
			3 6
			3 7
			4 4
			1 2 3 4
			1 2
			2 3
			3 4
			5 1
			1
			1 2
			1 3
			1 4
			1 5
			5 2
			4 5
			1 2
			2 3
			1 4
			4 5
			10 8
			1 2 3 4 5 8 9 10
			2 10
			10 5
			5 3
			3 1
			1 7
			7 4
			4 9
			8 9
			6 1
			10 9
			1 2 4 5 6 7 8 9 10
			1 3
			3 9
			9 4
			4 10
			10 6
			6 7
			7 2
			2 5
			5 8`,
			`2
			2
			0
			1
			4
			5`,
		},
		{
			`3
			6 1
			3
			1 2
			1 3
			3 4
			3 5
			2 6
			5 3
			1 2 5
			1 2
			1 3
			2 4
			3 5
			7 1
			2
			3 2
			2 6
			6 1
			5 6
			7 6
			4 5`,
			`0
			2
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1881F, testCases, 0)
}