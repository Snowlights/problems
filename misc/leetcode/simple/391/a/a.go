package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	cnt := 0
	for tmp := x; tmp > 0; tmp /= 10 {
		cnt += tmp % 10
	}
	if x%cnt == 0 {
		return cnt
	}
	return -1
}
