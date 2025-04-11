//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2050/E
// https://codeforces.com/problemset/status/2050/problem/E
func TestCF2050E(t *testing.T) {
	t.Log("Current test is [CF2050E]")
	testCases := [][2]string{
		{
			`7
			a
			b
			cb
			ab
			cd
			acbd
			ab
			ba
			aabb
			xxx
			yyy
			xyxyxy
			a
			bcd
			decf
			codes
			horse
			codeforces
			egg
			annie
			egaegaeg`,
			`1
			0
			2
			0
			3
			2
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2050E, testCases, 0)
}
