package _0

func sum(a, b int) int {
	return a + b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
