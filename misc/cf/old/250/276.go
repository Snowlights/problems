package _50

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
)

func CF276D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var left, right uint64
	fmt.Fscan(r, &left, &right)
	fmt.Fprintln(w, 1<<bits.Len64(left^right)-1)

}
