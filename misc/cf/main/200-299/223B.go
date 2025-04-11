//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF223B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	preCnt := make([]int, n)
	cnt := 0
	last := [26]int{}
	for i, b := range s {
		if cnt < m && t[cnt] == b {
			cnt++
			last[b-'a'] = cnt
		}
		if c := last[b-'a']; c > 0 {
			preCnt[i] = c
		}
	}

	cnt = m
	last = [26]int{}
	for i := n - 1; i >= 0; i-- {
		b := s[i]
		if cnt > 0 && t[cnt-1] == b {
			last[b-'a'] = cnt
			cnt--
		}
		sufCnt := int(1e9)
		if c := last[b-'a']; c > 0 {
			sufCnt = c
		}
		if preCnt[i] < sufCnt {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")

}

func main() { CF223B(os.Stdin, os.Stdout) }
