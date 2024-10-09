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
	Fscan(in, &s)

	palindrome := func(s string) bool {
		n := len(s)
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-1-i] {
				return false
			}
		}
		return true
	}
	if palindrome(s) &&
		palindrome(s[:(len(s)-1)/2]) &&
		palindrome(s[(len(s)+3)/2-1:]) {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}

}

func main() { run(os.Stdin, os.Stdout) }
