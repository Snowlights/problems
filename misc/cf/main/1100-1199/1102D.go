//go:build main
// +build main

package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

func CF1102D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s []byte
	Fscan(in, &n, &s)
	zero, one := bytes.Count(s, []byte("0")), bytes.Count(s, []byte("1"))
	two := n - zero - one

	for i, v := range s {
		switch v {
		case '1':
			if one > n/3 && zero < n/3 {
				s[i] = '0'
				one--
				zero++
			}
		case '2':
			if two > n/3 && zero < n/3 {
				s[i] = '0'
				two--
				zero++
				continue
			}
			if two > n/3 && one < n/3 {
				s[i] = '1'
				two--
				one++
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		switch s[i] {
		case '0':
			if two < n/3 && zero > n/3 {
				s[i] = '2'
				two++
				zero--
				continue
			}
			if one < n/3 && zero > n/3 {
				s[i] = '1'
				one++
				zero--
			}
		case '1':
			if one > n/3 && two < n/3 {
				s[i] = '2'
				one--
				two++
			}
		}
	}

	Fprintln(out, string(s))
}

func main() { CF1102D(os.Stdin, os.Stdout) }
