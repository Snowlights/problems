package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var low, high uint
	Fscan(in, &low, &high)
	if low == high {
		Print(1)
		return
	}
	ans := high - low + 1
	mask := uint(1)<<(bits.Len(high^low)-1) - 1
	high &= mask
	low &= mask
	nh := bits.Len(high)
	if bits.Len(low) <= nh {
		ans += mask - high
	} else {
		ans += mask - low + 1<<nh - high
	}
	Fprint(out, ans)

}

func main() { run(os.Stdin, os.Stdout) }
