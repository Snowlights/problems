package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func mergeSort(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := append(a[:0:0], a[:n/2]...)
	right := append(a[:0:0], a[n/2:]...)
	cnt := mergeSort(left) + mergeSort(right)
	l, r := 0, 0
	for i := range a {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			cnt += len(right) - r
			a[i] = left[l]
			l++
		} else {
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] += a[i-1] - k
	}
	Fprint(out, mergeSort(a))

}

func main() { run(os.Stdin, os.Stdout) }
