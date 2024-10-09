package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	var k int
	Fscan(in, &s, &k)

	type pair struct {
		i, pre int
	}
	h := make(map[pair]int)
	var dfs func(i, pre int, limit, num bool) int
	dfs = func(i, pre int, limit, num bool) int {
		if i == len(s) {
			if pre <= k && num {
				return 1
			}
			return 0
		}
		res := 0
		if !limit && num {
			dv, ok := h[pair{i, pre}]
			if ok {
				return dv
			}
			defer func() {
				h[pair{i, pre}] = res
			}()
		}

		if !num {
			res += dfs(i+1, pre, false, false)
		}

		down, up := 0, 9
		if !num {
			down = 1
		}
		if limit {
			up = int(s[i] - '0')
		}
		for ; down <= up; down++ {
			res += dfs(i+1, pre*down, limit && down == up, true)
		}
		return res
	}

	Fprintln(out, dfs(0, 1, true, false))

}

func main() { run(os.Stdin, os.Stdout) }
