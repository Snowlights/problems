//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 给定N个数.
// 问有多少连续子段满足，这一段内的数中，出现次数最多的一个数的出现次数>=K.

// 固定右端点，右移左端点，标记区间内个数大于k的值，累加答案

func CF190D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}

	ans, l := 0, 0
	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
		for cnt[v] >= k {
			cnt[arr[l]]--
			l++
		}
		ans += l
	}

	Fprintln(out, ans)
}

func main() { CF190D(os.Stdin, os.Stdout) }
