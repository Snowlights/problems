package _100_1200

// 1106
func parseBoolExpr(expression string) bool {
	stack := []byte{}
	for i := range expression {
		v := expression[i]
		switch v {
		case ')':
			f, t, ans := 0, 0, 'f'
			for stack[len(stack)-1] == 'f' || stack[len(stack)-1] == 't' || stack[len(stack)-1] == ',' {
				switch stack[len(stack)-1] {
				case 'f':
					f++
				case 't':
					t++
				}
				stack = stack[:len(stack)-1]
			}
			op := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch op {
			case '!':
				if f > 0 {
					ans = 't'
				}
			case '&':
				if f == 0 {
					ans = 't'
				}
			case '|':
				if t > 0 {
					ans = 't'
				}
			}
			stack = append(stack, byte(ans))
		default:
			stack = append(stack, v)
		}
	}

	return stack[0] == 't'
}
