package main

func countSeniors(details []string) (ans int) {
	for _, s := range details {
		// 对于数字字符，&15 等价于 -'0'，但是不需要加括号
		if s[11]&15*10+s[12]&15 > 60 {
			ans++
		}
	}
	return
}
