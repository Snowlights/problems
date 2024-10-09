package _021autumn

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestA(t *testing.T) {
	// source = [[1,3],[5,4]], target = [[3,1],[6,5]]
	//t.Log("Current test is [a]")
	//examples := [][]string{
	//	{
	//		`[[1,3],[5,4]]`, `[[3,1],[6,5]]`,
	//		`1`,
	//	},
	//	{
	//		`[[1,2,3],[3,4,5]]`, `[[1,3,5],[2,3,4]]`,
	//		`0`,
	//	},
	//}
	//targetCaseNum := 0 // -1
	//if err := testutil.RunLeetCodeFuncWithExamples(t, minimumSwitchingTimes, examples, targetCaseNum); err != nil {
	//	t.Fatal(err)
	//}

	//t.Log("Current test is [a]")
	//exampleIns := [][]string{{`[[1,3],[5,4]]`, `[[3,1],[6,5]]`}, {`[[1,2,3],[3,4,5]]`, `[[1,3,5],[2,3,4]]`}}
	//exampleOuts := [][]string{{`1`}, {`0`}}
	//targetCaseNum := 0
	//if err := testutil.RunLeetCodeFuncWithCase(t, minimumSwitchingTimes, exampleIns, exampleOuts, targetCaseNum); err != nil {
	//	t.Fatal(err)
	//}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumSwitchingTimes, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestB(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,2,8,9]`, `3`,
			`18`,
		},
		{
			`[3,3,1]`, `1`,
			`0`,
		},
	}
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, maxmiumScore, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestC(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`["....X.","....X.","XOOO..","......","......"]`,
			`3`,
		},
		{
			`[".X.",".O.","XO."]`,
			`2`,
		},
		{
			`[".......",".......",".......","X......",".O.....","..O....","....OOX"]`,
			`4`,
		},
	}
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, flipChess, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestD(t *testing.T) {
	t.Log("Current test is [d]")
	var examples = [][]string{
		{
			`[[3,3,1],[3,2,1]]`,
			`[[4,3]]`,
			`2`,
			`1`,
		},
		{
			`[[1,3,2],[4,3,1],[7,1,2]]`,
			`[[1,0],[3,3]]`,
			`4`,
			`2`,
		},
		{
			`[[2,2,1]]`,
			`[[1,3],[1,4]]`,
			`3`,
			`1`,
		},
	}
	targetCaseNum := -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, circleGame, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestE(t *testing.T) {
	t.Log("Current test is [e]")
	examples := [][]string{
		{
			`["W","N","ES","W"]`,
			`2`,
		},
		{
			`["NS","WE","SE","EW"]`,
			`3`,
		},
	}
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, trafficCommand, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
