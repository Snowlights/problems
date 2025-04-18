// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/agc012/submit?taskScreenName=agc012_b
// 对拍：https://atcoder.jp/contests/agc012/submissions?f.LanguageName=Go&f.Status=AC&f.Task=agc012_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`7 7
			1 2
			1 3
			1 4
			4 5
			5 6
			5 7
			2 3
			2
			6 1 1
			1 2 2`,
			`2
			2
			2
			2
			2
			1
			0`,
		},
		{
			`14 10
			1 4
			5 7
			7 11
			4 10
			14 7
			14 3
			6 14
			8 11
			5 13
			8 3
			8
			8 6 2
			9 7 85
			6 9 3
			6 7 5
			10 3 1
			12 9 4
			9 6 6
			8 2 3`,
			`1
			0
			3
			1
			5
			5
			3
			3
			6
			1
			3
			4
			5
			3`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/agc012/tasks/agc012_b
