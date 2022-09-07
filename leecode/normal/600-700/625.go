package _00_700

import "math"

// 627
// UPDATE salary
// SET
//    sex = CASE sex
//        WHEN 'm' THEN 'f'
//        ELSE 'm'
//    END;
//

// 633
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}
