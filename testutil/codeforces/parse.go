package codeforces

import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"strings"
	"time"
)

var (
	// true: test only one case in AssertEqualRunResultsInf / CheckRunResultsInf
	Once bool

	// true: only print test case number when assertion failed
	DisableLogInput bool

	DebugTLE = 2 * time.Second

	AssertOutput = true
)

func IsDebugging() bool {
	pid := int32(os.Getppid())

	// We loop in case there were intermediary processes like the gopls language server.
	for pid != 0 {
		p, err := process.NewProcess(pid)
		if err != nil {
			return false
		}
		name, err := p.Name()
		if err != nil {
			return false
		}
		if strings.HasPrefix(name, "dlv") {
			return true
		}
		pid, _ = p.Ppid()
	}
	return false
}
