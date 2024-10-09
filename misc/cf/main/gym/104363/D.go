package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CFD(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { CFD(os.Stdin, os.Stdout) }
