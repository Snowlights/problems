package _50

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF691C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)
	if !strings.Contains(s, ".") {
		s += "." // 标准化：补上小数点
	}
	s = strings.Trim(s, "0") // 标准化：去掉前后缀的 0

	i := strings.Index(s, ".")
	e := i - strings.IndexAny(s, "123456789") // 计算指数
	if e > 0 {
		e--
	}
	s = strings.Trim(s[:i]+s[i+1:], "0") // 去掉小数点，再次去掉前后缀的 0，得到小数部分
	if len(s) > 1 {
		s = s[:1] + "." + s[1:]
	}
	fmt.Fprint(w, s)
	if e != 0 {
		fmt.Fprint(w, "e", e)
	}

}
