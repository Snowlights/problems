package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n)
	type pair struct {
		s   string
		val int
	}
	cnt, m := make([]*pair, 0, n), make(map[string]*pair)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		p, ok := m[s]
		if !ok {
			p = &pair{s: s}
			m[s] = p
			cnt = append(cnt, p)
		}
		p.val++
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i].val > cnt[j].val || cnt[i].val == cnt[j].val && cnt[i].s < cnt[j].s
	})

	for _, v := range cnt {
		if v.val == cnt[0].val {
			Fprintln(out, v.s)
		} else {
			break
		}
	}

}

func main() { run(os.Stdin, os.Stdout) }
