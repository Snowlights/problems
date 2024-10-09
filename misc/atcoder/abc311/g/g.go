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

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
		mn := append([]int{}, a[i]...)
		pre := make([]int, m+1)
		for up := i; up >= 0; up-- {
			s := 0
			for j, v := range a[up] {
				mn[j] = min(mn[j], v)
				s += v
				pre[j+1] += s
			}
			posL := make([]int, m)
			posR := make([]int, m)
			st := []int{-1}
			for j, v := range mn {
				for len(st) > 1 && v < mn[st[len(st)-1]] {
					posR[st[len(st)-1]] = j
					st = st[:len(st)-1]
				}
				posL[j] = st[len(st)-1]
				st = append(st, j)
			}
			for _, j := range st[1:] {
				posR[j] = m
			}
			for j, v := range mn {
				ans = max(ans, (pre[posR[j]]-pre[posL[j]+1])*v)
			}
		}
	}
	Fprint(out, ans)
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() { run(os.Stdin, os.Stdout) }
