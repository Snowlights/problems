// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/practice2/submit?taskScreenName=practice2_g
// 对拍：https://atcoder.jp/contests/practice2/submissions?f.LanguageName=Go&f.Status=AC&f.Task=practice2_g&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{
			`6 7
			1 4
			5 2
			3 0
			5 5
			4 1
			0 3
			4 2`,
			`4
			1 5
			2 4 1
			1 2
			2 3 0`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/practice2/tasks/practice2_g
