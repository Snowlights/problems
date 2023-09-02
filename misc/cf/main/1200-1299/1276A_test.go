//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1276/A
// https://codeforces.com/problemset/status/1276/problem/A
func TestCF1276A(t *testing.T) {
	t.Log("Current test is [CF1276A]")
	testCases := [][2]string{
		{
			`4
			onetwone
			testme
			oneoneone
			twotwo`,
			`2
			6 3
			0
			
			3
			4 1 7 
			2
			1 4`,
		},
		{
			`10
			onetwonetwooneooonetwooo
			two
			one
			twooooo
			ttttwo
			ttwwoo
			ooone
			onnne
			oneeeee
			oneeeeeeetwooooo`,
			`6
			18 11 12 1 6 21 
			1
			1 
			1
			3 
			1
			2 
			1
			6 
			0
			
			1
			4 
			0
			
			1
			1 
			2
			1 11 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1276A, testCases, 0)
}
