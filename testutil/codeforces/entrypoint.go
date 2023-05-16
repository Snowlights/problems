package codeforces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type ioFunc func(io.Reader, io.Writer)

func isTLE(f func()) bool {
	if DebugTLE == 0 || IsDebugging() {
		f()
		return false
	}

	done := make(chan struct{})
	timer := time.NewTimer(DebugTLE)
	defer timer.Stop()
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}

func AssertEqualCase(t *testing.T, runFunc ioFunc, rawText string, targetCaseNum int) {
	rawText = strings.TrimSpace(rawText)
	if rawText == "" {
		t.Fatal("rawText is empty")
	}

	sepInput := "inputCopy"
	if !strings.Contains(rawText, sepInput) {
		sepInput = "input"
	}
	sepOutput := "outputCopy"
	if !strings.Contains(rawText, sepOutput) {
		sepOutput = "output"
	}

	testCases := [][2]string{}
	examples := strings.Split(rawText, sepInput)
	for _, s := range examples {
		s = strings.TrimSpace(s)
		if s != "" {
			splits := strings.Split(s, sepOutput)
			testCases = append(testCases, [2]string{splits[0], splits[1]})
		}
	}

	AssertEqualStringCase(t, runFunc, testCases, targetCaseNum)
}

func AssertEqualStringCase(t *testing.T, runFunc ioFunc, testCases [][2]string, targetCaseNum int) {
	if len(testCases) == 0 {
		return
	}

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(testCases) + 1
	}

	allPassed := true
	for curCaseNum, tc := range testCases {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		input := removeExtraSpace(tc[0])
		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize { // 截断过长的输入
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		expectedOutput := removeExtraSpace(tc[1])

		mockReader := strings.NewReader(input)
		mockWriter := &strings.Builder{}
		_f := func() { runFunc(mockReader, mockWriter) }
		if targetCaseNum == 0 && isTLE(_f) {
			allPassed = false
			t.Errorf("Time Limit Exceeded %d\nInput:\n%s", curCaseNum+1, inputInfo)
			continue
		} else if targetCaseNum != 0 {
			_f()
		}
		actualOutput := removeExtraSpace(mockWriter.String())

		t.Run(fmt.Sprintf("Case %d", curCaseNum+1), func(t *testing.T) {
			if !assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d\nInput:\n%s", curCaseNum+1, inputInfo) {
				allPassed = false
				handleOutput(actualOutput)
			}
		})
	}

	// 若有测试用例未通过，则前面必然会打印一些信息，这里直接返回
	if !allPassed {
		return
	}

	// 若测试的是单个用例，则接着测试所有用例
	if targetCaseNum > 0 {
		t.Logf("case %d is passed", targetCaseNum)
		AssertEqualStringCase(t, runFunc, testCases, 0)
		return
	}

	t.Log("OK")
}

func AssertEqualFileCase(t *testing.T, runFunc ioFunc, dir string, targetCaseNum int) {
	AssertEqualFileCaseWithName(t, runFunc, dir, "in*.txt", "ans*.txt", targetCaseNum)
}

func AssertEqualFileCaseWithName(t *testing.T, runFunc ioFunc, dir, inName, ansName string, targetCaseNum int) {
	inputFilePaths, err := filepath.Glob(filepath.Join(dir, inName))
	if err != nil {
		t.Fatal(err)
	}
	answerFilePaths, err := filepath.Glob(filepath.Join(dir, ansName))
	if err != nil {
		t.Fatal(err)
	}
	if len(inputFilePaths) != len(answerFilePaths) {
		t.Fatal("missing sample files")
	}
	if len(inputFilePaths) == 0 {
		t.Log("[WARN] empty test file")
		return
	}

	testCases := make([][2]string, len(inputFilePaths))
	for i, path := range inputFilePaths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][0] = string(data)
	}
	for i, path := range answerFilePaths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][1] = string(data)
	}

	AssertEqualStringCase(t, runFunc, testCases, targetCaseNum)
}
