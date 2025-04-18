// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc379/submit?taskScreenName=abc379_f
// 对拍：https://atcoder.jp/contests/abc379/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc379_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`5 3
			2 1 4 3 5
			1 2
			3 5
			1 4`,
			`2
			0
			1`,
		},
		{
			`10 10
			2 1 5 3 4 6 9 8 7 10
			3 9
			2 5
			4 8
			5 6
			3 8
			2 10
			7 8
			6 7
			8 10
			4 10`,
			`1
			3
			1
			2
			1
			0
			1
			1
			0
			0`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/abc379/tasks/abc379_f
