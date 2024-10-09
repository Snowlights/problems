package _021spring

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestElectricCarPlan(t *testing.T) {
	examples := [][]string{
		{
			`[[1,3,3],[3,2,1],[2,1,3],[0,1,4],[3,0,5]]`,
			`6`, `1`, `0`, `[2,10,4,1]`,
			`43`,
		},
		{
			`[[0,4,2],[4,3,5],[3,0,5],[0,1,5],[3,2,4],[1,2,8]]`,
			`8`, `0`, `2`, `[4,1,1,3,2]`,
			`38`,
		},
	}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, electricCarPlan, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestMaxGroupNumber(t *testing.T) {
	examples := [][]string{
		{
			`[2,2,2,3,4]`,
			`1`,
		},
		{
			`[2,2,2,3,4,1,3]`,
			`2`,
		},
		{
			`[1,2,2,2,3,4,5]`,
			`2`,
		},
	}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, maxGroupNumber, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestMinRecSize(t *testing.T) {
	examples := [][]string{
		{
			`[[2,3],[3,0],[4,1]]`,
			`48.00000`,
		},
		{
			` [[1,1],[2,3]]`,
			`0.00000`,
		},
	}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, minRecSize, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestGuardCastle(t *testing.T) {
	examples := [][]string{
		{
			`["S.C.P#P.", ".....#.S"]`,
			`3`,
		},
		{
			`["SP#P..P#PC#.S", "..#P..P####.#"]`,
			`-1`,
		},
		{
			`["SP#.C.#PS", "P.#...#.P"]`,
			`0`,
		},
		{
			`["CP.#.P.", "...S..S"]`,
			`4`,
		},
	}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, guardCastle, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
