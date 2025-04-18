// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc107/submit?taskScreenName=arc107_d
// 对拍：https://atcoder.jp/contests/arc107/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc107_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`4 2`,
			`2`,
		},
		{
			`2525 425`,
			`687232272`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/arc107/tasks/arc107_d
