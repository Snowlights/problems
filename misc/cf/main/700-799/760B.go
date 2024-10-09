//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF760B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, tot, k int
	Fscan(in, &n, &tot, &k)
	tot -= n
	ans := 1 + sort.Search(tot, func(m int) bool {
		m++
		st := max(m-k+1, 0)
		cnt := (st + m) * (m - st + 1) / 2
		st = max(m-(n-k), 0)
		cnt += (st + m - 1) * (m - st) / 2
		return cnt > tot
	})
	Fprint(out, ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { CF760B(os.Stdin, os.Stdout) }
