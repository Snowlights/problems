// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc160/submit?taskScreenName=abc160_f
// 对拍：https://atcoder.jp/contests/abc160/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc160_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`3
			1 2
			1 3`,
			`2
			1
			1`,
		},
		{
			`2
			1 2`,
			`1
			1`,
		},
		{
			`5
			1 2
			2 3
			3 4
			3 5`,
			`2
			8
			12
			3
			3`,
		},
		{
			`8
			1 2
			2 3
			3 4
			3 5
			3 6
			6 7
			6 8`,
			`40
			280
			840
			120
			120
			504
			72
			72`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/abc160/tasks/abc160_f
