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

	var n int
	Fscan(in, &n)
	a, b := make([][]int, n), make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	for i := range b {
		b[i] = make([]int, n)
		for j := range b[i] {
			Fscan(in, &b[i][j])
		}
	}

	compare := func() bool {
		for i, row := range a {
			for j, v := range row {
				if v == 1 && b[i][j] != v {
					return false
				}
			}
		}
		return true
	}
	exchange := func() {
		newA := make([][]int, n)
		for i := range newA {
			newA[i] = make([]int, n)
			for j := range newA[i] {
				newA[i][j] = a[j][n-1-i]
			}
		}
		a = newA
	}
	if compare() {
		Fprintln(out, "Yes")
		return
	}
	exchange()
	if compare() {
		Fprintln(out, "Yes")
		return
	}
	exchange()
	if compare() {
		Fprintln(out, "Yes")
		return
	}
	exchange()
	if compare() {
		Fprintln(out, "Yes")
		return
	}
	Fprintln(out, "No")
}

func main() { run(os.Stdin, os.Stdout) }
